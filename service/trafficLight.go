package service

import "fmt"

// "context"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"


type TrafficLight struct {
	ID      int
	Address string
	Status  bool
}

// type TrafficLightService struct {
// 	db *mongo.Database
// }

var light1 = TrafficLight{ID: 1, Address: "123 Main St", Status: true}
var light2 = TrafficLight{ID: 2, Address: "456 Elm St", Status: true}
var light3 = TrafficLight{ID: 3, Address: "789 Oak St", Status: true}
var light4 = TrafficLight{ID: 4, Address: "101 Maple Ave", Status: true}
var light5 = TrafficLight{ID: 5, Address: "202 Pine St", Status: true}

var TrafficLights = []TrafficLight{light1, light2, light3, light4, light5}

func ListAllLights() []TrafficLight {
	return []TrafficLight{light1, light2, light3, light4, light5}
}

func ChangeLightStatus(light *TrafficLight) {
	light.Status = !light.Status
	fmt.Println("Status mudado com sucesso")
}

func ChangeAllStatus() {
	go ChangeLightStatus(&light1)
	go ChangeLightStatus(&light2)
	go ChangeLightStatus(&light3)
	go ChangeLightStatus(&light4)
	go ChangeLightStatus(&light5)
}

