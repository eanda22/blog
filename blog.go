package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
    "encoding/json"
    "io/ioutil"
    "net/http"
)


var all_blogs[] string

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blog")
	})
	// post request
	http.HandleFunc("/post", post_request)

	// opens the server at port 80
	http.ListenAndServe(":80", nil)
}

func get_request(w http.ResponseWriter, r *http.Request) {
	// gets the json data
	json_body, err := json.Marshal(all_blogs)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	w.Write(json_body)
	/*
	*  add_entry() must be called to create the database
	*  Not sure how to put this JSON data into the database
	*/
}

func post_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
	}
}

func add_entry(id int, title string, body string) {
	
	// creates a new database and creates a table with an id, title, and body
	database, _ := sql.Open("sqlite3", "./blogger.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS blog_entry (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	statement.Exec()

	//inserts the blog entry into the database
	statement, _ = database.Prepare("INSERT INTO blog_entry (id, title, body) VALUES (?, ?, ?)")
	statement.Exec(id, title, body)

}


