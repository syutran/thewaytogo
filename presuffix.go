//presuffix.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Dose the string \"%s\" have prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
