package main

import "fmt"

func main(){
	var worldPopulation = 8_000_000_000 
	var growthRate  = 9

	fmt.Println("YEAR     ANTICIPATED-P       NUMERICAL-INCREASE")
	for year := 1; year <= 75; year++ {
		var numericalIncrease  =  (growthRate / 100 * worldPopulation)
		worldPopulation = worldPopulation + numericalIncrease
		fmt.Println( year , "         "    , worldPopulation ,"              " ,  numericalIncrease)
	}
}
