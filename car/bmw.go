package car

import "strconv"

type BMW struct {}

func (c BMW)Drive() {
	println("bmw is driving")
}

func (c BMW)OpenDoor(number int) {
	println("bmw is opening door number " + strconv.Itoa(number))
}