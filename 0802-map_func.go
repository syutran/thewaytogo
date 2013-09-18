//0802-map_func.go

package main

import (
	"fmt"
)

func main() {
	fm := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(fm)
}
