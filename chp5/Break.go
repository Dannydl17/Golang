package main

import "fmt"

func main(){
	var count int
	for count := 1; count <= 10; count++ {
		 if(count == 5){
			 return
		 }
		fmt.Printf("%d ", count)
	}
	 fmt.Println("%nBroke out of loop at count = %d%n", count)
 
}