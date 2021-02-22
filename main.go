package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct{
	ID string `json:"id"`
	Title string `json:"title"`
}

var db *sql.DB
var err error

func main()  {
	db, err := sql.Open("mysql", "root:Nomzano100%@tcp(127.0.0.1:3306)/rest-api-go")

	if err != nil {
        panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPosts).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	var posts []Post

	result, err := db.Query("SELECT id, title FROM posts")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next(){
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func createPosts(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
}

func getPost(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
}

func updatePost(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
}

func deletePost(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
}