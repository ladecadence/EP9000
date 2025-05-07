package main

import (
	"fmt"

	ep9000 "github.com/ladecadence/EP9000.git/pkg/EP9000"
)

func main() {
	data := make(chan []uint8)
	scanner, err := ep9000.New("/dev/ttyACM0", 115200)
	if err != nil {
		panic(err)
	}
	go scanner.Listen(data)

	select {
	case recv := <-data:
		fmt.Printf("Data: %s\n", recv)
	}
}
