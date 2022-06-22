package main

import (
	"fmt"
	"os"
)

func main() {
	// deferS are not run on Exit
	defer fmt.Println("!")
	os.Exit(100)
}