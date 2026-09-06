package main

import "fmt"

var age, income int
var hasBadCredit bool

func main() {
	fmt.Print("enter your age: ")
	fmt.Scan(&age)
	fmt.Print("enter your income: ")
	fmt.Print("are you have bad credit history?(true/false) ")
	fmt.Scan(&hasBadCredit)
	if age < 18 {
		fmt.Println("Loan rejected: Underage")
	} else if hasBadCredit {
		fmt.Println("Loan rejected: Bad credit history")
	} else if income < 10_000_000 {
		fmt.Println("Loan rejected: Low income")
	} else {
		fmt.Println("Loan approved")
	}
}
