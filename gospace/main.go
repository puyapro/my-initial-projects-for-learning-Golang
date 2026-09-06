package main

import "fmt"

func main() {
	var earthWeight int
	var gravityRatio float64
	fmt.Print("enter earth weight: ")
	fmt.Scan(&earthWeight)
	fmt.Print("enter gravity ratio of your planet: ")
	fmt.Scan(&gravityRatio)
	fmt.Println(float64(earthWeight) * gravityRatio)
}
