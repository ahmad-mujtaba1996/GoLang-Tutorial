package main

// maps -> hash, object, dictionary
func main() {

	// creating map

	// m := make(map[string]string) // -> [string] is key type, string is value type

	// setting an element
	// m["name"] = "Ahmad"
	// m["area"] = "backend"

	// get an element
	//fmt.Println(m["name"], m["area"])
	// fmt.Println(m["phone"])	// if key not found in map then it returns zero value

	// m := make(map[string]int)

	// m["age"] = 30
	// m["price"] = 1000
	// fmt.Println(m["phone"])
	// fmt.Println(len(m)) // length of map

	// deleting an element
	// delete(m, "price")
	// fmt.Println(m)

	// clearing a map
	// clear(m)
	// fmt.Println(m)

	// Alternative way to create map
	// m := map[string]int{"age":30, "price":1000}
	// fmt.Println(m)

	// multiple return values are supported in Go
	// k, ok := m["price"] // k -> value, ok -> boolean indicating if key exists

	// fmt.Println(k)

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// Comparison in Maps
	// m1 := map[string]int{"age": 30, "price": 1000}
	// m2 := map[string]int{"age": 30, "price": 1000}

	// fmt.Println(maps.Equal(m1, m2))
}
