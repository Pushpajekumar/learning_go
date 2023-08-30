package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in go lang")

	// no inheritance in go lang; np super or parent



	dhruv := User{"Dhruv", "XXXXXXXXXXXX", true, 22}

	fmt.Println(dhruv)

	fmt.Printf("Dhruv details are: %+v\n", dhruv)

	fmt.Printf("Name is %v and email is %v\n", dhruv.Name, dhruv.Email)
	dhruv.GetStatus()
	dhruv.NewMail()
	fmt.Printf("Name is %v and email is %v\n", dhruv.Name, dhruv.Email)

}

type User struct {
	Name   string
	Email  string
	Status	bool
	Age int
}

func (u User) GetStatus() string {
	fmt.Println("user is ", u.Status)
	return ""
}

func (u User) NewMail() {
	u.Email = "test@go.dev"

	fmt.Println("new mail ", u.Email)
}