package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
        "encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

type blog struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var all_blogs[] string

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blog")
	})
	// post request
	http.HandleFunc("/post", post_request)

	// opens the server at port 8080
	http.ListenAndServe(":8080", nil)
}

func get_request(w http.ResponseWriter, r *http.Request) {
	// reads out the request into a slice of bytes
	post, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	// appends the post to the existing array of posts (as strings) and prints them all
	all_blogs = append(all_blogs, string(post))
	for i := 0; i < len(all_blogs); i++ {
		fmt.Fprintf(w, all_blogs[i])
	}
	/*
	 * Not sure how to take the slice of strings and put it into the database
	 */
}

func post_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var post blog
		err := decoder.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(post)
		fmt.Fprintf(w, "Post Saved")
	}
}

func add_entry(title string, body string) {
	
	// creates a new database and creates a table with an id, title, and body
	database, _ := sql.Open("sqlite3", "./blogger.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS blog_entry (title TEXT, body TEXT)")
	statement.Exec()

	//inserts the blog entry into the database
	statement, _ = database.Prepare("INSERT INTO blog_entry (title, body) VALUES (?, ?)")
	statement.Exec(title, body)

}


