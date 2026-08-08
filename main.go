package main

import (
	datastructures "dsa-go/data-structures"
	"fmt"
)

func main() {
	arr := datastructures.NewDynamicArray[int](4)
	fmt.Println(arr)
}
