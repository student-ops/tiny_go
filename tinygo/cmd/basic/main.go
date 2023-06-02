package main

import (
	"machine"
	"time"
	"strconv"
)
func initUART() { 
	machine.UART2.Configure(machine.UARTConfig{}) 
} 
 
func putchar(c []byte) { 
	machine.UART2.Write(c) 
} 

func main() {
	count := 0
	initUART()
	for {
		count++;
		putchar([]byte("Uprint \"hello " +strconv.Itoa(count)+"\"" ))
		// fmt.Printf("hello world %d",count)
		println(count)
		time.Sleep(time.Second)
	}
}