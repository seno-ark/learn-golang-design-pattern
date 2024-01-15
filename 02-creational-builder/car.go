package main

// A Builder of type car
type CarBuilder struct {
	v VehicleProduct
}

func NewCarBuilder() IBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) SetStructure() IBuilder {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) SetWheels() IBuilder {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() IBuilder {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
