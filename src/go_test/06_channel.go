package main

import (
	"fmt"
)

func main() {
	pipline := make(chan bool)
	close(pipline)
	fmt.Println(<-pipline)
}
