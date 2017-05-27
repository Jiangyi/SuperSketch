package auth

import (
	"net/http"
	"os"
	"github.com/gorilla/securecookie"
	"log"
)


var(
	hashKey = []byte(os.Getenv("SECURE_COOKIE_HASHKEY"))
	blockKey = []byte(os.Getenv("SECURE_COOKIE_BLOCKKEY"))
	sc = securecookie.New(hashKey, blockKey)
)

func DecodeSessionID(cookie *http.Cookie) (string, error) {
	return decodeCookie(SESSION_ID, cookie)
}

func decodeCookie(cookieKey string, cookie *http.Cookie) (string, error) {
	var cookieVal string
	err := sc.Decode(cookieKey, cookie.Value, &cookieVal)
	if err != nil {
		return "", err
	}
	log.Println("Decoded cookie is: " + cookieVal)
	return cookieVal, nil

}
