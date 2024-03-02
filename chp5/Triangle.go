package main

import "fmt"

func main(){
	fmt.Print("Enter a number:  ")
    var number int
	fmt.Scanln(&number)
	for row := 0; row <number; row++ {
		for column := 0; column <=row; column++ {
			fmt.Print("*" + " ")
		}
		for column := 0; column <number-row; column++ {
			fmt.Print(" " + " ")
		}
		for column := 0; column <number-row; column++ {
			fmt.Print("*" + " ")
		}
		for column := 0; column <=row; column++ {
			fmt.Print(" " + " ")
		}
		for column := 0; column <=row; column++ {
			fmt.Print(" " + " ")
		}
		for column := 0; column <number-row; column++ {
			fmt.Print("*" + " ")
		}
		for column := 0; column <number-row; column++ {
			fmt.Print(" " + " ")
		}
		for column := 0; column <row + 1; column++ {
			fmt.Print("*" + " ")
		}
		fmt.Println();
	}
}