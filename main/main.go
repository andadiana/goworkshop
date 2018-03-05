package main

import (
	"fmt"
	"goworkshop/model"
	"goworkshop/web"
	"goworkshop/persistence"
	"goworkshop/importer"
)

func main() {

	persistence.InitDB()
	model.Authors = importer.ImportAuthors()
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	model.Books = importer.ImportBooks()
	fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}
