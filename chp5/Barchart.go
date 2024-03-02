package main

import "fmt"

func main(){
	

	var firstNumber = 0
	var secondNumber = 0
	var thirdNumber = 0
	var fourthNumber = 0
	var fifthNumber = 0
	var number = 0
	for count := 1; count <= 5; count++ {
		fmt.Print("Enter number between 1 and 30: ")
		fmt.Scanln(&number)
		if number < 1 || number > 30 {
			fmt.Println("Invalid number.Please try again")
			count--;
			continue
		}
		if count == 1 {
			firstNumber = number
		}
		if count == 2 {
			secondNumber = number
		}
		if count == 3 {
			thirdNumber = number
		}
		if count == 4 {
			fourthNumber = number
		}
		if count == 5 {
			fifthNumber = number
		}
	}
	fmt.Println(firstNumber)
	fmt.Println(secondNumber)
	fmt.Println(thirdNumber)
	fmt.Println(fourthNumber)
	fmt.Println(fifthNumber)


	for number := 1; number<=firstNumber; number++ {
		fmt.Print("*")
	}
	fmt.Println()

	for number := 1; number<=secondNumber; number++ {
		fmt.Print("*")
	}
      fmt.Println()

	for number := 1; number<=thirdNumber; number++ {
		fmt.Print("*")
	}
	fmt.Println()

	for number := 1; number<=fourthNumber; number++ {
		fmt.Print("*")
	}
	fmt.Println()

	for number := 1; number<=fifthNumber; number++ {
		fmt.Print("*")
	}
	fmt.Println();

}