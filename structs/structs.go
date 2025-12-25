// Go does not have classes, but it has structs which can be used to create complex data types.
package main

import (
	"fmt"
	"time"
)

// define a struct
type Customer struct {
	phone string
	name  string
	email string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Coming from time package. Its precision is in nanoseconds.

	//Struct embedding
	customer Customer
}

// Constructors are not present in Go Language. So need to create a hack
// In Go, Naming convention usually starts from capital alphabet.
func createOrder(id string, amount float32, status string) *Order {
	creatOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &creatOrder
}

// To connect this function with the struct, we need to use the receiver type that is (order Order)
func (order *Order) changeStatus(status string) { // while changing any property of Order struct, make sure to add pointer * against reciever type in order to correctly modify struct property value
	order.status = status // dereferencing operator not required as struct automatically handles it for us.

}

func (order Order) getOrderAmount() float32 { // Can remove * operator while fetching values for modification * operator is required in function
	return order.amount
}

func main() {
	// One way to initialize a struct
	//    var order Order = Order{
	// 	id: "1234567890",
	// 	amount: 100.00,
	// 	status: "pending",
	// 	createdAt: time.Now(),
	//    }

	// Another way to initialize a struct
	myOrder := Order{
		id:     "1234567890",
		amount: 100.00,
		status: "pending",
	}

	myOrder.createdAt = time.Now().UTC().Truncate(time.Millisecond)

	fmt.Println("Initial Order:", myOrder)

	newOrder := Order{
		id:        "2",
		amount:    100.00,
		status:    "recieved",
		createdAt: time.Now(),
	}

	fmt.Println("New Order:", newOrder)

	myOrder.status = "shipped"
	fmt.Println("Order Status:", myOrder.status)

	myOrder.changeStatus("Recieved")
	fmt.Println("Order Status Changed:", myOrder.status)

	fetchedAmount := myOrder.getOrderAmount()
	fmt.Println("Fetched Result: ", fetchedAmount)

	//If you don't set any field then default field value is set to zero
	// int => 0, float => 0, string => "", boolean => false
	currentOrder := Order{
		id: "1234567890",
		// amount: 100.00,
		status: "pending",
	}

	fmt.Println("Current Order:", currentOrder)

	//Creating order using modified constructor hack function
	secondaryOrder := createOrder("784", 300, "In Process")
	fmt.Println("Secondary Order is: ", secondaryOrder)
	fmt.Println("Secondary Order Amount is: ", secondaryOrder.amount)

	//Sometimes we need to have one object related to one usecase so don't need to create struct blueprint like done earlier with "type Order struct"
	// inline struct
	language := struct {
		name   string
		isGood bool
	}{"goLang", true}

	fmt.Println(language)

	nextCustomerOrder := Order{
		id:        "1234567890",
		amount:    100.00,
		status:    "pending",
		createdAt: time.Now(),
		customer: Customer{
			phone: "1234567890",
			name:  "John Doe",
		},
	}

	nextCustomerOrder.customer.name = "Syed Ahmad Mujtaba"
	fmt.Println("Next Customer Order:", nextCustomerOrder)
}
