package main

// Product
type VehicleProduct struct {
	Structure string
	Wheels    int
	Seats     int
}

// Builder interface
type IBuilder interface {
	SetStructure() IBuilder
	SetWheels() IBuilder
	SetSeats() IBuilder
	GetVehicle() VehicleProduct
}

func getBuilder(builderType string) IBuilder {
	if builderType == "car" {
		return NewCarBuilder()
	}

	if builderType == "motorbike" {
		return NewBikeBuilder()
	}
	return nil
}

// Director
type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) Construct() {
	d.builder.SetStructure().SetWheels().SetSeats()
}

func (d *Director) GetVehicle() VehicleProduct {
	return d.builder.GetVehicle()
}
