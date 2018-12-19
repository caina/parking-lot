package car

type Car struct {
	Plate string
	Color string
}

func (car Car) ToString() string {
	return "[" + car.Plate + " : " + car.Color + "]"
}
