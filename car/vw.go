package car

import "strconv"

type VW struct {}

func (c VW)Drive() {
	println("vw is driving")
}

func (c VW)OpenDoor(number int) {
	println("vw is opening door number " + strconv.Itoa(number))
}