package main

import (
	"fmt"
	"time"
)

// Написать собственную функцию Sleep.

func Sleep(i int) {
	<-time.After(time.Second * time.Duration(i))
}

func main() {
	fmt.Println("start")
	Sleep(2)
	fmt.Println("stop")
}
