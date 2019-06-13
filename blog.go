package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
    	"encoding/json"
	"net/http"
	"log"
)

type blog struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var all_blogs[] blog

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blog")
	})
	// post request
	http.HandleFunc("/post", post_request)

	// opens the server at port 8080
	http.ListenAndServe(":8080", nil)
}

func get_request(w http.ResponseWriter, r *http.Request) []blog {
	// returns the json array of blogs
	return all_blogs
}

func post_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var post blog
		err := decoder.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		// adds the new blog entry to the database
		add_entry(post.Title, post.Body)
		// adds the new blog entry to the array of json objects
		all_blogs = append(all_blogs, post)
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





