package main

import (
	"machine"
	"time"
	"fmt"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	count := 0
	for {
		count++;
		fmt.Printf("hello world %d",count)
		time.Sleep(time.Second)
		led.Low()
		time.Sleep(time.Millisecond * 500)
		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}