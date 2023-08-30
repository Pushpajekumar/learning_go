package main

import "fmt"

func main() {
	fmt.Println("Welcome to the maps in go lang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"
	languages["go"] = "Go"

	fmt.Println("List of all Languages: ", languages)

	fmt.Println("js shorts for: ", languages["js"])

	delete(languages, "rb")

	fmt.Println("List of all Languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}