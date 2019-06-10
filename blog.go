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

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blog")
	})
	// post request
	http.HandleFunc("/post", post_request)

	// get request
	http.HandleFunc("/get", get_request)

	// delete request
	http.HandleFunc("/delete/", delete_request)

	//put request
	http.HandleFunc("/put/", put_request)

	// opens the server at port 8080
	http.ListenAndServe(":8080", nil)
}

func get_request(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all_blogs)
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
		add_entry(post.Id, post.Title, post.Body)

		// adds the new blog entry to the array of json objects
		all_blogs = append(all_blogs, post)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
		log.Println("Post Saved")
	}
}

func delete_request(w http.ResponseWriter, r *http.Request) {
       	if r.Method == "DELETE" {
		// takes the id from the url
		id := strings.TrimPrefix(r.URL.Path, "/delete/")
		for index, item := range all_blogs {
			if string(item.Id) == id {
				// deletes post 
				all_blogs = append(all_blogs[:index], all_blogs[index + 1:]...)

				// deletes post from database
				delete_entry(all_blogs[index].Id, all_blogs[index].Title, all_blogs[index].Body)
	
				// returns the new json array 
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(all_blogs)
				log.Println("Post Deleted")
				return
			}
		}
	}

}

func put_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		// takes the id from the url
		id := strings.TrimPrefix(r.URL.Path, "/put/")
		for index, item := range all_blogs {
			if string(item.Id) == id {
				// deletes the original post
				all_blogs = append(all_blogs[:index], all_blogs[index + 1:]...)
				decoder := json.NewDecoder(r.Body)
				var post blog
				err := decoder.Decode(&post)
				if err != nil {
					log.Fatal(err)
				}
				// deletes the old post and adds the updates post to the database
				delete_entry(all_blogs[index].Id, all_blogs[index].Title, all_blogs[index].Body)
				add_entry(post.Id, post.Title, post.Body)

				// adds the updated post in place of the original and returns json value
				all_blogs = append(all_blogs, post)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(post)
				log.Println("Post Updated")
				return
			}
		}
	}
}

func add_entry(id string, title string, body string) {
	// creates a new database and creates a table with an id, title, and body
	database, _ := sql.Open("sqlite3", "./blogger.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS blog_entry (id TEXT, title TEXT, body TEXT)")
	statement.Exec()

	//inserts the blog entry into the database
	statement, _ = database.Prepare("INSERT INTO blog_entry (id, title, body) VALUES (?, ?, ?)")
	statement.Exec(id, title, body)
}

// Implement
func delete_entry(id string, title string, body string) {
		// creates a new database and creates a table with an id, title, and body
		database, _ := sql.Open("sqlite3", "./blogger.db")
		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS blog_entry (id TEXT, title TEXT, body TEXT)")
		statement.Exec()

		statement, _ = database.Prepare("DELETE FROM blog_entry WHERE id=" + id)
		statement.Exec(id, title, body)
}



