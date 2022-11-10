package main

import (
	"database/sql"
	"io"
	"fmt"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test1234"
	dbname   = "database"
)

func main() {

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	CheckError(err)

	fmt.Println(string(b))


	//////////


	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "Jokes"("Test") values(string(b))`
	_, e := db.Exec(insertStmt)
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}