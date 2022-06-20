package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cannot work with 42")
	}
	return arg + 100, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("Cannot work with %d - %s\n", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "42 is verboten"}
	}
	return arg + 10, nil
}

func main() {
	if res, err := f1(10); err != nil {
		fmt.Println("Error with 10")
	} else {
		fmt.Println(res)
	}

	if res, err := f1(42); err != nil {
		fmt.Println("Error with 42")
	} else {
		fmt.Println(res)
	}

	res1, err1 := f2(100)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res2, err2 := f2(42)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}

}