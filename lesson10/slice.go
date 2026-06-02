package main

import "fmt"

func task1() []string {
	fmt.Println("Enter all the cities")
	cities := []string{}
	for i := 0; i < 3; i++ {
		var temp string
		fmt.Scan(&temp)
		cities = append(cities, temp)
	}
	return cities
}

func task2() []string {
	cities := []string{"Delhi", "Ghaziabad"}
	cities = append(cities, "Basti")
	cities = append(cities, "Lucknow")
	cities = append(cities, "Gorakhpur")
	return cities
}

func task3() int {
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func task4() ([]string, []int) {
	student := []string{"Student1", "Student2", "Student3", "Student4", "Student5"}
	marks := []int{100, 30, 60, 70, 90}
	return student, marks
}

func main() {
	mycities := task1()
	mcities := task2()
	sum := task3()
	student, marks := task4()
	fmt.Println(mycities, "\n", mcities, "\n", sum)
	for i, mark := range marks {
		fmt.Println(student[i], ":", mark)
	}
}
