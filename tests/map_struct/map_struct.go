package main

import "fmt"

type User1 struct {
	Id int
	UserName string
}
func main() {
	mapa := make(map[int]User1)
	mapa[1] = User1{Id:1,UserName:"隔壁张三"}
	mapa[2] = User1{Id:2,UserName:"QQ飞车"}
	mapa[3] = User1{
		Id:       3,
		UserName: "嘎嘎嘎",
	}
	fmt.Println(mapa)
}
