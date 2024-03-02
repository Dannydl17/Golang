package main

import "fmt"

func main(){
	for number := 0; number < 30; number++ {
		if (number % 3 == 0) {
			fmt.Println(number)
		}
	}
}