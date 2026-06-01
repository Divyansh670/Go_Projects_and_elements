package main
import "fmt"

func main(){
fmt.Println("Enter Your Name")
var name string
fmt.Scan(&name)
fmt.Println("Enter Your Age")
var age int
fmt.Scan(&age)
fmt.Println("Name and Age is",name,age)
}

