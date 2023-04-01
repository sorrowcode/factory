package main

import c "github.com/sorrowcode/factory/car"

func main() {
	car := c.ProduceCar(c.CarType(1))
	car.Drive()
}
