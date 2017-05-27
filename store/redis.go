package store

import (
	"os"
	"log"
	"github.com/garyburd/redigo/redis"
)

var (
	Redis redis.Conn
)

func Init() {
	conn, err := redis.Dial("tcp", ":6379"); if err != nil {
		log.Fatal("Failed to get Redis connection!")
		os.Exit(-42)
	}
	Redis = conn
}
