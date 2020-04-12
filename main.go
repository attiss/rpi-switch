package main

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

var (
	pins = []rpio.Pin{
		rpio.Pin(2),
		rpio.Pin(3),
		rpio.Pin(4),
		rpio.Pin(17),
		rpio.Pin(27),
		rpio.Pin(22),
		rpio.Pin(10),
		rpio.Pin(9),
	}
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	for _, pin := range pins {
		pin.Output()
		pin.High()
	}

	for _, pin := range pins {
		if os.Args[1] == "on" {
			pin.Low()
		} else {
			pin.High()
		}
	}
}
