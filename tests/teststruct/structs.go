package main

import "fmt"

type User struct {
	Id int
	UserName string
}

func main() {
	var user1 User = User{
		Id:       0,
		UserName: "123",
	}
	fmt.Println(user1)

	var user2 User
	user2.Id = 1
	user2.UserName = "124"
	fmt.Println(user2)

	user3 := User{3,"125"}
	fmt.Println(user3)
}