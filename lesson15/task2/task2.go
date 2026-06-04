package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "abc"
	num, err := strconv.Atoi(str)
	fmt.Printf("err: %v\n", err)
	fmt.Println(num)
}
