package auth

import (
	"net/http"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/kvnxiao/supersketch/store"
	"github.com/kvnxiao/supersketch/db"
	"fmt"
	"strings"
	"time"
	"strconv"
)

const(
	SESSION_ID string = "Session-ID"
	FLASH_MSG string = "Flash-Msg"
	FIRST_NAME string = "FirstName"
	LAST_NAME string = "LastName"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// TODO: Sanitize all teh things
	username := strings.ToLower(r.Form["username"][0])
	email := r.Form["email"][0]
	password := r.Form["password"][0]
	firstName := r.Form["firstName"][0]
	lastName := r.Form["lastName"][0]

	fmt.Printf("Username: %s; email: %s; password: %s; firstName: %s; lastName: %s", username, email, password, firstName, lastName)
	if db.UserExists(username) {
		sendErr(w, "Already registered. Why you do dis.", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); {
		if err != nil {
			sendErr(w, "Error in bcrypt hashing. Please contact the administrator.", http.StatusBadRequest)
			return
		}
	}

	err = db.AddUser(username, email, hashedPassword, firstName, lastName)
	if err != nil {
		SendFlashMessage(w, r, err.Error(), "/register")
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully Registered as: " + username))
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := strings.ToLower(r.Form["username"][0])
	password := r.Form["password"][0]
	invalidCredentialsMsg := "Invalid credentials entered. Please try again."

	userID, passHash, firstName, lastName, err := db.GetUserByUsername(username)
	if err != nil {
		SendFlashMessage(w, r, invalidCredentialsMsg, "/login")
		return
	}

	err = bcrypt.CompareHashAndPassword(passHash, []byte(password))
	if err != nil {
		SendFlashMessage(w, r, invalidCredentialsMsg, "/login")
		return
	}

	cookieKey := "COOKIE:" + strconv.Itoa(userID)
	log.Println(cookieKey)

	_, err = store.Redis.Do("HMSET", cookieKey, FIRST_NAME, firstName, LAST_NAME, lastName); if err != nil {
		sendErr(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = store.Redis.Do("EXPIRE", cookieKey, 604800); if err != nil {
		sendErr(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if encoded, err := sc.Encode(SESSION_ID, cookieKey); err == nil {
		sessionIDCookie := &http.Cookie{
			Name: SESSION_ID,
			Value: encoded,
			Path: "/",
			Secure: false, // FIXME
			Expires: time.Now().Add(time.Second * 604800),
		}
		http.SetCookie(w, sessionIDCookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		sendErr(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func SendFlashMessage(w http.ResponseWriter, r *http.Request, msg string, redirectPath string) {
	flashCookie := &http.Cookie{
		Name: FLASH_MSG,
		Value: msg,
		Secure: false,
		HttpOnly: false,
		Expires: time.Now().Add(time.Second * 5),
	}
	http.SetCookie(w, flashCookie)
	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func Play(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SESSION_ID); if err == nil {
		sessionID, err := DecodeSessionID(cookie); if err != nil {
			SendFlashMessage(w, r, err.Error(), "/login")
		}
		result, err := store.Redis.Do("HMGET", sessionID, FIRST_NAME, LAST_NAME); if err != nil {
			sendErr(w, "Redis lookup failed", http.StatusInternalServerError)
		}
		if result != nil {
			switch value := result.(type) {
			case []interface{}:
				log.Printf("First name: %s; Last name: %s\n", value[0], value[1])
				break;
			default:
				log.Print("Unknown type")
				break;
			}
			http.ServeFile(w, r, "view/dist/index.html")
		}

	} else {
		SendFlashMessage(w, r, err.Error(), "/login")
	}
}

func sendErr(w http.ResponseWriter, err string, httpErrorCode int) {
	log.Println(err)
	w.WriteHeader(httpErrorCode)
	w.Write([]byte(err))
}