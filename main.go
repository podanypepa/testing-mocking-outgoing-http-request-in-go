package main

import (
	"fmt"
	"log"
)

func main() {
	apiurl := "https://jsonplaceholder.typicode.com/posts"

	data, err := fetchData(apiurl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
