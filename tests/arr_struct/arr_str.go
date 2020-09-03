package main

import "fmt"

type User struct {
	Id int
	Name string
}

func main() {
	user := User{Id:1,Name:"hello"}
	fmt.Println(user)

	arr := [4]int{1,2,3,4}
	fmt.Println(arr)

	arr2 := [4]User{
		{Id:1,Name:"hello"},
		{Id:2,Name:"hello2"},
		{Id:3,Name:"hello3"},
		{Id:4,Name:"hello"},
	}
	fmt.Println(arr2)
}
