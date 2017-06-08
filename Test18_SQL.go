package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	id       int
	name     string
	password string
	email    string
}

func (u User) String() string {
	return fmt.Sprintf("user(id = %v, name = %v, password = %v, email = %v)",
		u.id, u.name, u.password, u.email)
}

func insert(db *sql.DB) {
	Hash := md5.New()
	p := hex.EncodeToString(Hash.Sum([]byte("asdf")))
	SQL := fmt.Sprintf("INSERT INTO test_user (`id`, `name`, `password`, `email`) "+
		"VALUES (2, 'user01', '%s', 'abcd@qq.com')", p)

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

func query(db *sql.DB) User{
	var u User
	r := db.QueryRow("SELECT id,name,password,email FROM test_user WHERE id=1")
	var (
		id int
		name string
		password string
		email string
	)
	r.Scan(&id, &name, &password, &email)
	u = User{id, name, password, email}
	return u
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	insert(db)
}

func selectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(query(db).String()))
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
	http.HandleFunc("/query", selectHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
