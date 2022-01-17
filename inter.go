package main

import "fmt"

type Kitchen interface {
	state()
}

func beingontheplate(kitchen Kitchen) {
	kitchen.state()
}

type Stand struct{}
type Lies struct{}

func (s Stand) state() {
	fmt.Println("Предмет стоїть на столі")
}
func (l Lies) state() {
	fmt.Println("Предмет лежить на столі")
}

func main() {

	spoon := Lies{}
	plate := Stand{}
	beingontheplate(spoon)
	beingontheplate(plate)
}
