package main

import "goworkshop/test"
import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/model"
	"goworkshop/web"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println("Catalin")
	fmt.Println("Hi, my name is Alex")
	fmt.Println("Hi, my name is Tibi")
	fmt.Println("Hi, my name is Anda! :)")

	// var s = test.createSquare(11)
	// s.Color = "green"
	// fmt.Println(s)

	// var s = test.Square{
	// 	Length: 10,
	// }
	// s.Color = "green"

	var s = test.CreateSquare(10)
	fmt.Println(s.Area())


	//import books
	booksFileContent, err := ioutil.ReadFile("model/books.json")
	if err != nil {
		//fmt.Println("Error occurred")
		//os.Exit(1)
		//this is equivalent to:
		panic(err)
	}

	//fmt.Println(string(fileContent))

	if err = json.Unmarshal(booksFileContent, &model.Books); err != nil {
		panic(err)
	}

	fmt.Println("The deserialized data:")
	fmt.Println(model.Books)

	serializedData, err := json.Marshal(model.Books)
	if err != nil {
		panic(err)
	}
	fmt.Println("The serialized data is:")
	fmt.Println(string(serializedData))


	//import authors
	authorsFileContent, err := ioutil.ReadFile("model/authors.json")
	if err != nil {
		//fmt.Println("Error occurred")
		//os.Exit(1)
		//this is equivalent to:
		panic(err)
	}

	//fmt.Println(string(fileContent))

	if err = json.Unmarshal(authorsFileContent, &model.Authors); err != nil {
		panic(err)
	}

	fmt.Println("The deserialized data:")
	fmt.Println(model.Authors)

	serializedData, err = json.Marshal(model.Authors)
	if err != nil {
		panic(err)
	}
	fmt.Println("The serialized data is:")
	fmt.Println(string(serializedData))

	web.StartServer()
}
