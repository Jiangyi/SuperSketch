package db

import (
	"os"
	"log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	postgresDB *sql.DB
)

func Init() {
	db, err := sql.Open("postgres", "dbname=supersketch sslmode=disable")
	checkErr(err, -43)

	postgresDB = db
}

func UserExists(user string) (bool) {
	var id int
	Row := postgresDB.QueryRow("SELECT id FROM users WHERE username = $1", user);
	if Row == nil {
		return false
	} else {
		Row.Scan(&id)
		fmt.Printf("User ID: %d\n", id)
		return id != 0
	}

}

func GetUserByID(id int) (username string, firstName string, lastName string, err error) {
	err = postgresDB.QueryRow("SELECT username, firstName, lastName FROM users WHERE id = $1", id).Scan(&username, &firstName, &lastName)

	if err != nil {
		log.Print(err)
		return "", "", "", err
	}

	fmt.Printf("User ID: %d; Username: %s; firstName: %s; lastName: %s", id, username, firstName, lastName)

	return username, firstName, lastName, nil
}


func GetUserByUsername(username string) (id int, passHash []byte, firstName string, lastName string, err error) {
	err = postgresDB.QueryRow("SELECT id, bCryptPassHash, firstName, lastName FROM users WHERE username = $1", username).Scan(&id, &passHash, &firstName, &lastName)

	if err != nil {
		log.Print(err)
		return 0, nil, "", "", err
	}

	fmt.Printf("User ID: %d; Username: %s; firstName: %s; lastName: %s", id, username, firstName, lastName)

	return id, passHash, firstName, lastName, nil
}

func AddUser(username string, email string, bCryptHash []byte, firstName string, lastName string) (error) {
	var id int
	err := postgresDB.QueryRow("INSERT INTO users(username, email, bCryptPassHash, firstName, lastName) VALUES($1,$2,$3,$4,$5) returning id;", username, email, bCryptHash, firstName, lastName).Scan(&id)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func checkErr(err error, code int) {
	if err != nil {
		log.Fatal(err)
		os.Exit(code)
	}
}
