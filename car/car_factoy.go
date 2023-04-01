package car

func ProduceCar(carType CarType) Car {
	switch carType {
	case bmw:
		return BMW{}
	case mercedes:
		return Mercedes{}
	default:
		return VW{}
	}
}