//ifelse.go
package main

import "fmt"

func main() {
	var first int = 10
	var cond int
	if first <= 0 {
		fmt.Println("First is less than or equal to 0")

	} else if first > 0 && first < 5 {
		fmt.Println("first is between 0 and 5.")
	} else {
		fmt.Println("first is 5 or greater.")
	}

	if cond = 5; cond > 10 {
		fmt.Println("cond is greater than 10.")
	} else {
		fmt.Println("cond is not greater than 10.")
	}
}
