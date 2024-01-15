package main

import "fmt"

func main() {
	carBuilder := getBuilder("car")
	director := NewDirector(carBuilder)
	director.Construct()

	car := director.GetVehicle()
	fmt.Printf(">> car product %+v\n", car)

	motorbikeBuilder := getBuilder("motorbike")
	director.SetBuilder(motorbikeBuilder)
	director.Construct()

	motorbike := director.GetVehicle()
	fmt.Printf(">> bike product %+v\n", motorbike)
}
