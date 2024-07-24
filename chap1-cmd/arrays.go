package main

import ("fmt")

func array_learn() {

	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"a", "b", "c"}
	fmt.Println(arr)

	arr[1] = "b1"
	fmt.Println(arr)
}
