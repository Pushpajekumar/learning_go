package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcom to files in golang")

	content := "This is needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
		}

	fmt.Println("Length is: ", length)
	defer file.Close()
	
	readFile("./mylcogofile.txt")
}

func readFile (fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is: ", string(databyte))
}