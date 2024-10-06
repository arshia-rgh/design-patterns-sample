package creational

import (
	"errors"
)

type ICar interface {
	setName(name string)
	getName() string
	setSpeed(speed int)
	getSpeed() int
}

type Car struct {
	name  string
	speed int
}

func (car *Car) setName(name string) { car.name = name }
func (car *Car) getName() string     { return car.name }
func (car *Car) setSpeed(speed int)  { car.speed = speed }
func (car *Car) getSpeed() int       { return car.speed }

type BMW struct {
	Car
}

func newBMW() ICar {
	return &BMW{Car{
		name:  "BMW",
		speed: 300,
	}}
}

type Benz struct {
	Car
}

func newBenz() ICar {
	return &Benz{Car{
		name:  "Benz",
		speed: 250,
	}}
}

// factory

func getCar(carType string) (ICar, error) {
	if carType == "BMW" {
		return newBMW(), nil
	}

	if carType == "Benz" {
		return newBenz(), nil
	}

	return nil, errors.New("wrong car type")
}

//func main() {
//	bmw, _ := getCar("BMW")
//	benz, _ := getCar("Benz")
//
//	fmt.Println(bmw.getName(), bmw.getSpeed())
//	fmt.Println(benz.getName(), benz.getSpeed())
//
//}
