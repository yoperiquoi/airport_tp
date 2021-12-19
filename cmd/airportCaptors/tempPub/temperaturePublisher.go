package main

import (
	pubconfig "airport_tp/infernal/config/captorConfig"
	pubutils "airport_tp/infernal/utils/captorUtils"
	"fmt"
)

func main() {
	config := pubconfig.LoadConfig("temperature")
	fmt.Println(config)
	publisher := pubutils.Connect("tcp://localhost:1883", "publisher")

	publisher.Disconnect(250)
}
