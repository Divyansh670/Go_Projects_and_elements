package main

import "fmt"

func marks(n int) (string, int) {
	var name string
	var mark int
	fmt.Print("Enter student ", n, " name:")
	fmt.Scanln(&name)
	fmt.Print("Enter student ", n, " marks:")
	fmt.Scanln(&mark)
	return name, mark
}

func main() {
	var size int
	fmt.Println("Enter the number of students:")
	fmt.Scanln(&size)
	subjects := [3]string{"Maths", "Physics", "Chemistry"}
	fmt.Println("Subjects are:", subjects)
	marksheet := make([]int, size)
	students := make([]string, size)
	for i := 1; i <= len(students); i++ {
		students[i-1], marksheet[i-1] = marks(i)
		fmt.Println("Student", i, "name and marks are", students[i-1], marksheet[i-1])
	}

}
