package main

// A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

func NewBikeBuilder() IBuilder {
	return &BikeBuilder{}
}

func (b *BikeBuilder) SetWheels() IBuilder {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() IBuilder {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() IBuilder {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
