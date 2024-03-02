package main

import "fmt"

func main() {
	fmt.Println("Enter an integer:   ") 
	var number int 
	fmt.Scanln(&number)
	var largestNumber int
	largestNumber = number

	for counter := 1; counter < 10; counter++ {
		println("Enter a number: ")
		fmt.Scanln(&number)

		if number > largestNumber {
			largestNumber = number
		}
	}

	fmt.Println("LargestNumber is   ", largestNumber)
}
