package main

import "fmt"

func main() {
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println(a)
	var b [5]int
	b[0] = 1
	b[1] = 2
	b[2] = 3
	fmt.Println(b)
	c := [5]int{1,2,3,4,5}
	fmt.Println(c)
	d := [...]int{1,2,3,4,5,6}
	fmt.Println(d)
	//第一种for循环
	for i:=0;i<len(d) ;i++  {
		fmt.Println("key:",i,"value",d[i])
	}
	//第二种循环
	for k,v := range d {
		fmt.Println(k,v)
	}
	//第三种循环
	count := 0
	for {
		fmt.Println(1)
		count ++
		if count == 10 {
			break
		}
	}
}
