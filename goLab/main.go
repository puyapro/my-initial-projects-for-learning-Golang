package main

import "fmt"

var i string
var ii, iii float64

func main() {
	for {
		fmt.Scan(&i)
		switch i {
		case "exit":
			fmt.Println("Lab console terminated")
			return
		case "+":
			fmt.Scan(&ii, &iii)
			fmt.Printf("Result: %.4f\n", ii+iii)
		case "-":
			fmt.Scan(&ii, &iii)
			fmt.Printf("Result: %.4f\n", ii-iii)
		case "*":
			fmt.Scan(&ii, &iii)
			fmt.Printf("Result: %.4f\n", ii*iii)
		case "/":
			fmt.Scan(&ii, &iii)
			if iii == 0 {
				fmt.Println("Error: Invalid mixture")
			} else {
				fmt.Printf("Result: %.4f\n", ii/iii)
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}

