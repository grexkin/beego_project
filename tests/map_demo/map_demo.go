package main

import "fmt"

func main() {
	//map: 键值对
	var map_a map[string]string
	map_a = make(map[string]string)   //必须要这样初始化
	map_a["hello"] = "world"
	map_a["name"] = "albe"
	fmt.Println(map_a)
	//第二种方法
	var map_b map[string]string = make(map[string]string)
	map_b["hell"] = "ddd"
	fmt.Println(map_b)
	//第三种方法
	map_c := map[string]string{"name":"hello"}
	fmt.Println(map_c)
	//获取key对应的value
	age_value,ok := map_c["age"]
	if !ok {
		fmt.Println("age_value 不存在")
	}else {
		fmt.Println(age_value)   //key存在情况输出
	}
	map_c["age"] = "20"
	//循环获取map数据
	for i:=0;i<len(map_c);i++ {
		fmt.Println(i)  //只能获取index
	}
	//正确获取map数据的方式
	for k,v := range map_c {
		fmt.Println(k,v)
	}
}
