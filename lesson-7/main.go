package main

import "fmt"

// 配列型

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// arr4 := [...]string{"C", "D"}
	arr4 := [...]string{"C"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr1[0])

	arr1[1] = 10
	fmt.Println(arr1)

	arr2[2] = "C"
	fmt.Println(arr2)

	fmt.Println(len(arr2))
}
