package main

import "fmt"

type order struct {
	id   int
	name string
}

func (o *order) changeName(name string){ // pointer receiver need to modify the original struct
	o.name = name
}

func (o order) getName() string{  //  pointer receiver not needed to read the original struct
	return o.name
}
func NewOrder(id int,name string) *order {
	 myOrder := order{
		id:   id,
		name: name,
	}
	return &myOrder
}

func main() {
	myOrder := order{
		id:   1,
		name: "Order1",
	}
	myOrder.id=2
	myOrder.changeName("UpdatedOrder1")
	myOrder2:=NewOrder(2,"Order2")
	fmt.Println(myOrder2.id)
	config:=struct{
		name string
		age  int
	}{"Aniket",20}
	fmt.Println(config.name, config.age)
}
