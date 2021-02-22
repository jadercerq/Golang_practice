package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	channel := make(chan int)

	for i := 0; i < 5; i++ {
		canal := "Canal-" + strconv.Itoa(i)
		go worker(channel, canal)
	}

	for i := 0; i < 25; i++ {
		channel <- i
	}

}

func worker(channel chan int, canal string) {
	for i := range channel {
		fmt.Println(canal, i)
		time.Sleep(time.Second * 2)
	}
}
