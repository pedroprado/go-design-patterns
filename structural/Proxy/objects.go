package main

import "fmt"

type Drivable interface {
	Drive()
}

type Car struct {
	Driver *Driver
}

func NewCar(driver *Driver) *Car {
	return &Car{Driver: driver}
}

func (ref *Car) Drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

func NewDriver(age int) *Driver {
	return &Driver{Age: age}
}

type CarProxy struct {
	car *Car
}

func NewCarProxy(drive *Driver) *CarProxy {
	return &CarProxy{
		car: &Car{Driver: drive},
	}
}

func (ref *CarProxy) Drive() {
	if ref.car.Driver.Age < 16 {
		fmt.Println("Can't drive")
	} else {
		ref.car.Drive()
	}
}
