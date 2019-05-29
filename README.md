# blog

## Data Base 

The data base is created in the blog.go file using sqlite. The database has 3 columns. 
The id (which is the primary key), the title of the blog entry, and the body of the entry. 

## Web Server

A local webserver is created with 2 requests, GET and POST. POST should print all of the entries as JSON objects.
GET takes the title and body from the request and put it into the database.
