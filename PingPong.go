package main

import (
	"fmt"
	"time"
)

func Ping(c1 chan string) {
	for i := 0; ; i++ {
		c1 <- "Ping"
	}
}
func Pong(c2 chan string) {
	for i := 0; ; i++ {
		c2 <- "Pong"
	}
}
func imprimir(c1,c2 chan string) {
	for {
		msg := <-c1 
		msg2 := <-c2
		fmt.Println(msg)		
		time.Sleep(time.Second * 1)
		fmt.Println(msg2)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c1 chan  string = make(chan string) 
	var c2 chan  string = make(chan string) 
	
	go Ping(c1)
	go Pong(c2)
	go imprimir(c1, c2)
	
	
	var entrada string
	var entrada2 string
	fmt.Scanln(&entrada, &entrada2)
}

