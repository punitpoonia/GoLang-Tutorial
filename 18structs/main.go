package main

import (
	"fmt"
	"time"
)

// this is the syntax to create the struct use type keyword then name of the strcut and the the struct keyword

type order struct {
	id          int
	description string
	createdAt   time.Time
	status      string
	customer
}

//embeeding of structs one struct inside other

type customer struct {
	name  string
	phone string
}

// how to use constructor in golang

func newOrder(id int, description string, status string) *order {
	myOrder := order{
		id:          id,
		description: description,
		status:      status,
	}

	return &myOrder
}

// To modify the values in the struct we have to use pointer
// Struct automatically derefferenced the values no need to do it manually

func (o *order) changeStatus(status string) {

	o.status = status

}

func main() {
	// if you don't fill any values in struct by default it is zero
	// int > 0 , flaot > 0 ,boolean = flase , string = empty string
	myOrder := newOrder(1, "iphone14", "processing")
	// myOrder.id = 12345
	// myOrder.description = "Iphone 16Pro 256GB"
	// myOrder.createdAt = time.Now()
	// myOrder.status = "Completed"
	// myOrder.changeStatus("Delivered")

	// create a method for customer

	myCustomer := customer{
		name:  "Punit",
		phone: "8765932123",
	}

	myEmbeddedOrder := order{
		id:          1,
		description: "MacBook Air",
		status:      "In Transit",
		customer:    myCustomer, // this is how we assign values to embeded struct

		// *****another way******
		// customer:customer{name:  "Punit",phone: "8765932123",},
	}
	myEmbeddedOrder.customer.name = "Poonia" // thats how we can change the name or values in the struct
	fmt.Println(myOrder)
	fmt.Println(myEmbeddedOrder)

}
