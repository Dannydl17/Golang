package main

import "fmt"

func main(){
	for row := 1; row>0; row-- {
		for column :=0; column <=7; column++ {
			fmt.Print(" ")
		}
		for column := 0; column<1; column++ {
			fmt.Print(" *")
		}
		fmt.Println()

		for column :=0; column<=7; column++ {
			fmt.Print(" ")
		}
		for  inColumn :=0; inColumn<3; inColumn++ {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column<=6;column++ {
			fmt.Print(" ")
		}
		for column :=0; column<5; column++ {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column<=5;column++ {
			fmt.Print(" ")
		}
		for column :=0; column<7; column++ {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column<=4;column++ {
			fmt.Print(" ")
		}
		for column :=0; column<9; column++ {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column<=5;column++ {
			fmt.Print(" ")
		}
		for column :=7; column>0; column-- {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column<=6;column++ {
			fmt.Print(" ")
		}
		for column :=5; column>0; column-- {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column<=7; column++ {
			fmt.Print(" ")
		}
		for inColumn :=3; inColumn>0; inColumn-- {
			fmt.Print("*")
		}
		fmt.Println()

		for column :=0; column <=7; column++ {
		      fmt.Print(" ")
		}
		for column := 1; column>0; column-- {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}