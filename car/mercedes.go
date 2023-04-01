package car

import "strconv"

type Mercedes struct {}

func (c Mercedes)Drive() {
	println("mercedes is driving")
}

func (c Mercedes)OpenDoor(number int) {
	println("mercedes is opening door number " + strconv.Itoa(number))
}