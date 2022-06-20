package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	switch n % 2 {
	case 0:
		fmt.Println(n, "is even")
	case 1:
		fmt.Println(n, "is odd")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}

	whatAmI := func(a any) {
		switch a.(type) {
		case bool:
			fmt.Println("A boolean")
		case int:
			fmt.Println("An int")
		default:
			fmt.Println("Some other type")
		}
	}

	whatAmI("Hello")
	whatAmI(12345)
	whatAmI(12 > 2)

}