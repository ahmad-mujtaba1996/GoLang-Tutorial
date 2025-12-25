package main

func main() {
	// Switch statement example

	// i := 4
	// Simple Switch statement
	// switch i {
	// case 1:
	// 	println("One")
	// 	// break isn't needed in Go because it automatically handles that stuff
	// case 2:
	// 	println("Two")
	// case 3:
	// 	println("Four")
	// case 4:
	// 	println("Four")
	// default:
	// 	println("Other number")
	// }

	//multiple condition in switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("It's the weekend!")
	// default:
	// 	println("It's a weekday.")
	// }

	// type switch
	// function is first-class citizen in Go
	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("I'm an integer")
	// 	case string:
	// 		fmt.Println("I'm a string")
	// 	case bool:
	// 		fmt.Println("I'm a boolean")
	// 	case float64:
	// 		fmt.Println("I'm a float64")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }

	// whoAmI(43)
	// whoAmI("hello")
	// whoAmI(true)
	// whoAmI(3.14)
}
