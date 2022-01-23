package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*----int----*/
	var a int32 = 123
	fmt.Printf("%s\n", "hello,world")
	fmt.Printf("type:", reflect.TypeOf(a))
	fmt.Printf("%s\n", "")

	/*---string---*/
	s := "Hello"
	m := "World"
	sm := s + m
	fmt.Printf("%s\n", sm)

	multiNoStr := `Hello,
        World`
	fmt.Printf("%s\n", multiNoStr)

	/*--array,slice--*/
	arr := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var aSlice, bSlice []byte
	aSlice = arr[3:7]
	fmt.Printf("%s\n", aSlice)
	bSlice = aSlice[1:3]
	fmt.Printf("%s\n", bSlice)
	bSlice = aSlice[0:6]
	fmt.Printf("%s\n", aSlice[5:6])
	arr[3] = 'o'
	fmt.Printf("%s\n", string(bSlice[0]))

	/*-- map ------*/
	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map 有两个返回值，第二个返回值，如果不存在 key，那么 ok 为 false，如果存在 ok 为 true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素
	fmt.Printf("result:", string(Max(2, 1)))

}
