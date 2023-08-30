package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests in go lang")

	respone, err :=http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The respone is of type %T\n", respone)

 defer	respone.Body.Close() // closing the respone body is very important

 dataBytes, err := ioutil.ReadAll(respone.Body)

 if err != nil {
	panic(err)
 };

 content := string(dataBytes)

 fmt.Println(content)
}