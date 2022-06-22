package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.235", 10)
	fmt.Println(f)

	i, _ := strconv.ParseInt("12345", 10, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("12345")
	fmt.Println(k)

	_, e := strconv.ParseInt("hello", 0, 64)
	fmt.Println(e)
}