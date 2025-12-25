package main

import "fmt"

// enumerated types
type OrderStatus int

const (
	Pending OrderStatus = iota // iota is a predefined identifier in Go that is used to generate a sequence of numbers starting from 0. It continues to increment with addition in constants.
	Processing
	Shipped
	Delivered
)

// If enum set as int, to convert to string need below kind of functionality
func (s OrderStatus) String() string {
	// return []string{"Pending", "Processing", "Shipped", "Delivered"}[s]

	// Slice of strings corresponding to the iota indices
	names := []string{"Pending", "Processing", "Shipped", "Delivered"}

	// Safety check to prevent "out of bounds" errors
	if s < 0 || int(s) >= len(names) {
		return "Unknown"
	}
	return names[s]
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to: ", status)
	// fmt.Println("Order status changed to: ", status.String())
}

type PaymentMode string

const (
	Stripe PaymentMode = "Stripe"
	PayPal PaymentMode = "PayPal"
	Cash   PaymentMode = "Cash"
)

func changePaymentStatus(status PaymentMode) {
	fmt.Println("Payment status changed to: ", status)
}
func main() {
	changeOrderStatus(Shipped)
	changePaymentStatus(PayPal)
}
