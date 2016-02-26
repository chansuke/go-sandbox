package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Your faviorite number is", math.Nextafter(2, 10))
}
