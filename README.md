# golang workshop

# Part 1 - Write REST services with data loaded from files
### 1. hello world in GO
1. create main directory
1. create main.go file
1. declare package main
1. declare main function
1. import fmt package
1. make a call to fmt.Println and print "Hello world, my name is <your_name_here>"
### 2. defining the data model
1. create datamodel.go file
1. create AuthorDto struct with the following fields:
 - UUID string
 - FirstName string
 - LastName string
 - Birthday string
 - Death string
1. create BookDto struct with the following fields:
- UUID string
- Title string
- NoPages int
- ReleaseDate string
- Author AuthorDto
4. create a slice of AuthorDto which will hold all authors in the system
5. create a slice of BookDto which will hold all books in the system
### 3. add JSON marshalling
### 4. read sample data from files
### 5. start a web server, listening on a configured port
### 6. writing a simple REST endpoint
### 7. implement all REST endpoints with data loaded from files

# Part 2 - Add unit tests for all REST endpoints and create the persistence layer
### 8. write unit tests for all REST endpoints
### 9. add persistence mappings
1. delete importer package
1. remove importer references from main.go
1. remove importer references from book_handlers.go
1. remove importer references from author_handlers.go
1. rename BookDto to Book and AuthroDto to Author
1. create an Entity struct in datamodel.go file with UUID field
1. add `gorm:"primary_key"` primary key annotation to the Entity.UUID field
1. embed Entity struct in Book struct
1. add `gorm:"ForeignKey:AuthorUUID"` foreign key annotation to the Book.Author field
1. add `sql:"type:text REFERENCES author(uuid) ON DELETE CASCADE"` cascade delete annotation to the Book.AuthorUUID field
1. embed Entity struct in the Author struct
### 10. add persistence services to retrieve data from db
1. create datastore.go file in persistence package
1. create GormDataStore struct with a field named DBInstance of a type representing a pointer to a gorm.DB
1. create DataStore interface in datastore.go with all books and authors operations
1. create a Store variable of DataStore type
1. in persistence/config.go file initialize the store variable before returning the DBInstance
1. create books_datastore.go file in persistence package
1. create authors_datastore.go file in persistence package
1. implement DataStore interface in books_datastore.go and authors_datastore.go files by creating functions with a pointer receiver to a gorm.DB type
### 11. switch from loading data from db instead of files
1. in main.go init the persistence
1. in web/book_handlers remove Books type and books variable
1. in web/author_handlers remove Authors type and authors variable
1. use methods on persistence.Store to work with Book and Author entities

#Part 3 - Add error handling and logging
### 12 write unit tests for All REST endpoints using mocks for all external dependencies
