package main

import "fmt"

func main(){
	fmt.Print("Enter an integer:     ")
	var number int
	fmt.Scanf("%d", &number)

	var factorial = 1

	for counter := 0; counter < number; counter++{
		factorial = factorial * (number - counter);
	}

	fmt.Println("Factorial is " , factorial)
}