package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {

	fmt.Println("Building router..")

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/plot", Plot)
	router.GET("/split", Split)
	router.GET("/train", Train)
	router.GET("/neural", ExecuteNN)

	fmt.Println("Listening and Serving..")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func initDb() {
	var err error
	fmt.Println("initializing database")

	db, err = sql.Open("postgres", "host=localhost port=32772 user=admin password=admin dbname=postgres sslmode=disable")

	if err != nil {
		fmt.Println("error connection to db", err.Error())
	}
}
