package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcom to conditionals in golang")

	// loginCount := 23
	// var result string

	// if loginCount < 10 {
	// 	result = "Regular user"
	// }else if loginCount > 10 {
	// 	result = "Watch out"
	// }else {
	// 	result = "something else"
	// }

	// fmt.Println(result)

	// if num:= 3; num < 10 {
	// 	fmt.Println("Num is less than 10")
	// }

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can open")
	case 3:
		fmt.Println("Dice value is 3 and you can open")
	case 4:
		fmt.Println("Dice value is 4 and you can open")
	case 5:
		fmt.Println("Dice value is 5 and you can open")
	case 6:
		fmt.Println("Dice value is 6 and you can open")
	}
}