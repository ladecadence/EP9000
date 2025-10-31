package main

import (
	"fmt"

	ep9000 "github.com/ladecadence/EP9000"
)

func main() {
	data := make(chan []uint8)
	scanner, err := ep9000.New("/dev/ttyACM1", 115200)
	if err != nil {
		panic(err)
	}

	go func() {
		err := scanner.Listen(data)
		if err != nil {
			panic(err)
		}
	}()

	for {
		select {
		case recv := <-data:
			fmt.Printf("Data: %s\n", recv)
		}
	}
}
