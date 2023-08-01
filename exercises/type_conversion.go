package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Type Conversion
	var integer16 int16 = 127
	var integer32 int32 = 32767
	fmt.Println(int32(integer16) + integer32)

	// _ means we are not going to use this variable
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
