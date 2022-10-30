package main

import (
	"flag"
	"time"

	"gitea.k8s.attiss.xyz/attiss/rpi-switch/config"
	"github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"
)

// var (
// 	pins = []rpio.Pin{
// 		rpio.Pin(2),  // 8
// 		rpio.Pin(3),  // 7
// 		rpio.Pin(4),  // 6
// 		rpio.Pin(17), // 5
// 		rpio.Pin(27), // 4
// 		rpio.Pin(22), // 3
// 		rpio.Pin(10), // 2
// 		rpio.Pin(9),  // 1
// 	}
// )

var (
	configFile = flag.String("config", "config.yaml", "Path for the YAML configuration file.")
)

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	config, err := config.ReadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// // Set pin to output mode
	// for _, pin := range pins {
	// 	pin.Output()
	// 	pin.High()
	// }

	// for _, pin := range pins {
	// 	if os.Args[1] == "on" {
	// 		pin.Low()
	// 	} else {
	// 		pin.High()
	// 	}
	// }

	pin := rpio.Pin(config.PinAssignment.Relay1Pin)
	pin.Output()
	pin.High()

	for i := 0; i < 10; i++ {
		logger.Info("toggled")
		time.Sleep(2 * time.Second)
		pin.Toggle()
	}
}
