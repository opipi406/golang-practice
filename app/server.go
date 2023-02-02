package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int
	name string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/insert", insertHandler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func connectionDatabase() (*sql.DB, error) {
	database := os.Getenv("MYSQL_DATABASE")
	password := os.Getenv("MYSQL_PASSWORD")

	// NOTE: Dockerを使っている場合、localhostではなくコンテナ名になるので要注意 → [db]:3306
	connectionInfo := fmt.Sprintf("user:%s@tcp(db:3306)/%s", password, database)
	db, err := sql.Open("mysql", connectionInfo)

	return db, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ハロー・ワールド")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectionDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintln(w, "%v", err)
		panic(err.Error())
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintln(w, user.id, user.name)
	}
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	name := string(r.URL.Query().Get("name"))
	if name == "" {
		fmt.Fprintln(w, "クエリパラメータに\"name\"を含めてください")
		return
	}

	db, err := connectionDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	in, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := in.Exec(name)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "insert successful (id = %d)", lastId)
}
