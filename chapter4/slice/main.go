package main

import "fmt"

func main() {
	var slice1 []int
	fmt.Printf("type:%T--------value:%v", slice1, slice1)
	if slice1 == nil {
		fmt.Println("equal nil")
	}
	slice2 := []int{}
	if slice2 != nil {
		fmt.Println("not equal nil")
	}
}
