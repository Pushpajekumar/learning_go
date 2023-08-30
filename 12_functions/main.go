package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions in golang")
	greetings()

	result := adder(3, 5)
	fmt.Println("result of sum is ", result)

	proResult, myMessage := proAdder(3, 5, 7, 8)
	fmt.Println("pro result is ", proResult)
	fmt.Println("pro message is ", myMessage)

}

func adder(valOne int, valTwo int) int  {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0;

	for _, val := range values{
			total += val
	}
	return total, "hiiiiii"

}

func greetings() {  
	fmt.Println("Namastey from golang")
}