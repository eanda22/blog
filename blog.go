package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
    "encoding/json"
	"net/http"
	"log"
	"strings"

)

type blog struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var all_blogs[] blog
var database *sql.DB

func main() {
	var err error
	database, err = sql.Open("sqlite3", "./blogger.db")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS blog_entry (id TEXT, title TEXT, body TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blog")
	})

	http.HandleFunc("/post", post_request)
	http.HandleFunc("/get", get_request)
	http.HandleFunc("/delete/", delete_request)
	http.HandleFunc("/put/", put_request)
	
	http.ListenAndServe(":8080", nil)
}

func get_request(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("SELECT id, title, body FROM blog_entry")
	if err != nil {
		log.Fatal(err)
	}
	var blog_id, blog_title, blog_body string
	for rows.Next() {
		rows.Scan(&blog_id, &blog_title, &blog_body)
		temp_blog := blog{Id: blog_id, Title: blog_title, Body: blog_body }
		all_blogs = append(all_blogs, temp_blog)
	}
	rows.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all_blogs)
	all_blogs = nil
	
}

func post_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var post blog
		err := decoder.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		add_entry(post.Id, post.Title, post.Body)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
		log.Println("Post Saved")
	}
}

func delete_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := strings.TrimPrefix(r.URL.Path, "/delete/")

		delete_entry(id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(all_blogs)
		log.Println("Post Deleted")
		return
	}

}

func put_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := strings.TrimPrefix(r.URL.Path, "/put/")

		delete_entry(id)
		decoder := json.NewDecoder(r.Body)
		var post blog
		err := decoder.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		add_entry(post.Id, post.Title, post.Body)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
		log.Println("Post Updated")
		return
	}
}

func add_entry(id string, title string, body string) {
	statement, err := database.Prepare("INSERT INTO blog_entry (id, title, body) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(id, title, body)
	log.Println(id)
}

func delete_entry(blog_id string) {
		statement, err := database.Prepare("DELETE FROM blog_entry WHERE id=?")
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec(blog_id)
		log.Println("deleted")
}

