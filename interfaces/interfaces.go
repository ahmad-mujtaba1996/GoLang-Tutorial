package main

import "fmt"

type paymentGateway interface {
	pay(amount float32) // can put return type also like this -> pay(amount float32) (bool, error)
	sendMoney(amount float32)
}

// to make functions optional of any interface, here is the easy hack
// type childPaymentGateway struct{}

// func (childPaymentGateway) pay(amount float32)       {}
// func (childPaymentGateway) sendMoney(amount float32) {}

type Payment struct {
	paymentGateway
}

type JazzCash struct{}

type SadaPay struct{}

// Violates -> Open Closed Principle ->  because we are modifying the existing code by commenting JazzCash and adding SadaPay
func (p Payment) makePayment(amount float32) {
	// jazzCashPaymentGW := JazzCash{}
	// jazzCashPaymentGW.pay(amount)
	// sadaPayPaymentGW := SadaPay{}
	// sadaPayPaymentGW.pay(amount)
	p.pay(amount)
}

// In Go, we don't need to implement interface function explicitly. It's done implicitly. Go automatically understands the relevant interface by looking at the functions used and its signature. If same, then uses interfaces otherwise treats as normal function.
func (jazzCash JazzCash) pay(amount float32) {
	fmt.Println("making payment using jazzcash: ", amount)
}

func (jazzCash JazzCash) sendMoney(amount float32) {}

func (sadaPay SadaPay) pay(amount float32) {
	fmt.Println("making payment using sadapay: ", amount)
}

func (sadaPay SadaPay) sendMoney(amount float32) {}

func main() {
	newPayment := Payment{
		paymentGateway: JazzCash{},
	}
	newPayment.makePayment(100)
}

// Example for optional interface methods
// type Plugin interface {
//     Initialize()
//     OnEvent()
//     Cleanup()
// }

// // BasePlugin provides default empty methods
// type BasePlugin struct{}
// func (BasePlugin) Initialize() {}
// func (BasePlugin) OnEvent()    {}
// func (BasePlugin) Cleanup()    {}

// // MyPlugin only implements what it needs
// type MyPlugin struct {
//     BasePlugin
// }

// func (m MyPlugin) OnEvent() {
//     fmt.Println("Event received!")
// }
