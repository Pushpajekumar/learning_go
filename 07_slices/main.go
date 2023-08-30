package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" Welcom to video on slices")

	var fruitList = []string{"Apple", "Banana", "Orange"}

	fmt.Printf(" Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Tomato")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 200
	highScores[2] = 300
	highScores[3] = 400

	fmt.Println(highScores)

	highScores = append(highScores, 500, 600, 700)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove value from slice based on index

	var courses = []string{"Docker", "Kubernetes", "Go", "react", "angular", "swift"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)	

}