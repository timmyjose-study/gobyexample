package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{x: 1, y: -200}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("Type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 12345)
	fmt.Printf("float: %f\n", 123.45)
	fmt.Printf("float: %e\n", 123.45)
	fmt.Printf("float: %E\n", 123.45)
	fmt.Printf("char: %c\n", 123)
	fmt.Printf("bin: %b\n", 12345)
	fmt.Printf("hex: %x\n", 12345)

	fmt.Printf("string1: %s\n", "\"string\"")
	fmt.Printf("string2: %q\n", "\"string\"")
	fmt.Printf("string3: %x\n", "\"string\"")

	fmt.Printf("pointer: %p\n ", &p)

	fmt.Printf("|%10.3f|\n", 11/3.4)
	f := fmt.Sprintf("|%-10.3f|\n", 11/3.4)
	fmt.Println(f)

	fmt.Fprintf(os.Stderr, "io: %s\n", "some error")

}