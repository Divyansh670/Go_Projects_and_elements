package main

import (
	"golang-course/lesson14/greetings"
	"golang-course/lesson14/mathutils"
)

func main() {
	greetings.Sayhello()
	a := 23
	b := 24
	mathutils.Op(a, b)
}
