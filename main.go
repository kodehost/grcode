package main

import "fmt"

func main() {
	fmt.Println("Hi! Dont Go")

	pace := 10
	for i := 1; i < 10; i++ {
		pace += i
		fmt.Println("pace address:", &pace, "  /  Value:", pace)

	}

}
