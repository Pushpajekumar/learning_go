package main

import "fmt"

func main() {
	fmt.Println("Wlcome to array in go lang")

	var fruitList [5]string 

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))


	var vegList = [5]string{"Potato","Onion","Carrot","Cabbage","Cucumber"}
	fmt.Println(vegList)

}
