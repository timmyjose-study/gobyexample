package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // blocking
	fmt.Println("Timer1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer2 stopped")
	}

	// allow timer2 to run, if possible
	time.Sleep(2 * time.Second)
}