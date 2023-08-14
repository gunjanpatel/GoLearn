package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}

	// Best way to close file in GO
	// As defer make sure that it will be closed when execution is done
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

	newfile.Sync()
}
