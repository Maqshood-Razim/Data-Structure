package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("array task 1")

	v := time.Now()
	fmt.Println(v)

	fmt.Println("enter array limit")
	var size int
	fmt.Scan(&size)

	arr := make([]int, size)

	fmt.Println("enter array values")

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	found := false

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {

			if arr[i]+arr[j] == 10 {
				fmt.Printf("pair found %d + %d =10", arr[i], arr[j])
				found = true
			}

		}

	}

	if !found {
		fmt.Println("not found")
	}

	a := time.Now()

	fmt.Println(a)

	// //another method///

	// var limit int
	// fmt.Println("enter array limit")
	// fmt.Scan(&limit)

	// arr := make([]int, limit)
	// fmt.Println("enter array values")
	// for i := range arr {
	// 	fmt.Scan(&arr[i])
	// }

	// seen := make(map[int]bool)
	// for _, num := range arr {
	// 	if seen[10-num] {
	// 		fmt.Printf("pair found: %d + %d = 10\n", num, 10-num)
	// 		return
	// 	}
	// 	seen[num] = true
	// }

	// fmt.Println("not found")

	fmt.Println("enter array limit")
	var limit int
	fmt.Scan(&limit)

	arr1 := make([]int, limit)

	fmt.Println("enter array values")
	for i := 0; i < limit; i++ {
		fmt.Scan(&arr1[i])
	}

	last := limit - 1
	for i := 0; i <= last; i++ {
		if arr1[i] == 6 {

			arr1[i], arr1[last] = arr1[last], arr1[i]
			last--
			i--
		}
	}

	fmt.Println("new array")

	for i := 0; i < limit; i++ {
		fmt.Println(arr1[i])
	}
}