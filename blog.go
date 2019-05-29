package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

var all_blogs[] entry

// represents a blog entry
type entry struct {
	entry_id int
	entry_title string
	entry_body string
}

func add_entry(id int, title string, body string) {
	
	// creates a new database and creates a table with an id, title, and body
	database, _ := sql.Open("sqlite3", "./blogger.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS blog_entry (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	statement.Exec()

	//inserts the blog entry into the database
	statement, _ = database.Prepare("INSERT INTO blog_entry (id, title, body) VALUES (?, ?, ?)")
	statement.Exec(id, title, body)

	//inserts the blog entry into a slice of entries
	my_entry := entry{id, title, body}
	all_blogs = append(all_blogs, my_entry)
}

func list_entries() {
	var id int
	var title string
	var body string

	// iterates through the slice of entries and prints out the id, title, and body of each one
	for i := 0; i < len(all_blogs); i++ {
		id = all_blogs[i].entry_id
		title = all_blogs[i].entry_title
		body = all_blogs[i].entry_body

		fmt.Println("Blog Post #" + strconv.Itoa(id))
		fmt.Println("Title: " + title)
		fmt.Println("Body: " + body)
		fmt.Println(" ")
	}
}

