package main

import "fmt"

func main(){
	fmt.Print("Enter a number:  ") 
    var numberOfStar int 
	fmt.Scanln(&numberOfStar)
     
    for row := 1; row <= numberOfStar; row++ {
         for column := 0; column < row; column++ { 
            fmt.Print("*")
        }
        fmt.Println()
    }
}