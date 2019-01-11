package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	fmt.Println("Done listen port 8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func insertDb() {

}

// func queryDb() {
// 	// Execute the query
// 	results, err := db.Query("SELECT id, name FROM tags")
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}

// 	for results.Next() {
// 		var tag Tag
// 		// for each row, scan the result into our tag composite object
// 		err = results.Scan(&tag.ID, &tag.Name)
// 		if err != nil {
// 			panic(err.Error()) // proper error handling instead of panic in your app
// 		}
// 		// and then print out the tag's Name attribute
// 		log.Printf(tag.Name)
// 	}
// }
func connectDb() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:123456789@tcp(localhost:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // be careful deferring Queries if you are using transactions
	// defer insert.Close()

	// Execute the query
	// results, err := db.Query("SELECT id, name FROM test")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// log.Println(results)
	// for results.Next() {
	// 	var tag Tag
	// 	// for each row, scan the result into our tag composite object
	// 	err = results.Scan(&tag.ID, &tag.Name)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	// 	// and then print out the tag's Name attribute
	// 	log.Println(tag.Name)
	// }

	// var tag Tag
	// // Execute the query
	// err = db.QueryRow("SELECT id, name FROM test where id = ?", 2).Scan(&tag.ID, &tag.Name)
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// log.Println(tag.ID)
	// log.Println(tag.Name)
}

func main() {
	connectDb()
	handleRequests()
}
