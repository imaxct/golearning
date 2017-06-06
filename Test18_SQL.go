package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/hex"
	"crypto/md5"
	"net/http"
)

type User struct {
	id int
	name string
	password string
	email string
}

func insert(db *sql.DB) {
	Hash := md5.New()
	p := hex.EncodeToString(Hash.Sum([]byte("asdf")))
	SQL := fmt.Sprintf("INSERT INTO test_user (`id`, `name`, `password`, `email`) " +
			"VALUES (1, 'imaxct', '%s', 'abcd@qq.com')", p)

	res, err := db.Exec(SQL)
	if err != nil {
		log.Fatal(err)
	}
	maxId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id: %v, rows: %v\n", maxId, rows)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	insert(db)
}

var db *sql.DB
var err error

func main() {
	fmt.Println("Test")
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/insert", insertHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
