package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	fmt.Printf("err: %v\n", err)
	fmt.Println(num)
}
