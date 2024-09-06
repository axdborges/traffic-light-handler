package main

import (
	"fmt"
	"time"
	// "time"
	"trafficLights/service"
)

func main() {
	println("Listando todos os semáforos antes do paralelismo")
	fmt.Println(service.ListAllLights())
	time.Sleep(time.Second)

	service.ChangeAllStatus()
	
	time.Sleep(time.Second)
	println("Listando depois de mudar o status")
	fmt.Println(service.ListAllLights())
}