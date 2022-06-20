package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("Hello,", i)
		i++
	}

	sum := 0
	for j := 0; j < 10; j++ {
		sum += j
	}
	fmt.Println(sum)

	ctr := 0
	for {
		fmt.Println("ctr =", ctr)
		ctr++

		if ctr > 10 {
			break
		}
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}