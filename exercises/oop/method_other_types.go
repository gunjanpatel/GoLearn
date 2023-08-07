package main

import (
	"fmt"
	"strings"
)

type upperstring string

func (s upperstring) upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning")
	fmt.Println(s.upper())
}
