package main

import (
	"github.com/gorilla/context"
	"github.com/kvnxiao/supersketch/auth"
	"github.com/kvnxiao/supersketch/chat"
	"github.com/kvnxiao/supersketch/db"
	"github.com/kvnxiao/supersketch/store"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	r *chi.Mux
)

func setup() {
	db.Init()
	store.Init() // Init Redis
	// Create router
	r = chi.NewRouter()

	// Setup middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// Set up websocket hub
	hub := chat.NewHub()

	// Auth
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			serveTest(w, r)
		} else {
			auth.Register(w, r)
		}
	})

	// Auth
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			serveTest(w, r)
		} else {
			auth.Login(w, r)
		}
	})

	// Auth
	r.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		auth.Play(w, r)
	})
	// Endpoints
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "view/dist/static")
	r.FileServer("/static/", http.Dir(filesDir))
	r.Get("/", serveTest)
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Cookies())
		hub.HandleWs(w, r)
	})

	// DEBUG: check goroutine count
	//debugLog()
}

func serveTest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/dist/index.html")
}

func main() {
	setup()
	run()
	cleanup()
}

func cleanup() {
	store.Redis.Close()
}

func run() {
	http.ListenAndServe(":9001", context.ClearHandler(r))
}

func debugLog() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Printf("Goroutines: %d", runtime.NumGoroutine())
			}
		}
	}()
}
