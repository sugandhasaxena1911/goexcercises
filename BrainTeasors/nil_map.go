package main

import "fmt"

func main() {
	var m map[string]int //nil map
	fmt.Println(m)
	//access element
	fmt.Println(m["Hello"]) //value at key not there : so it prints corresponding zero_val

	//add elements
	//Panic will occur
	/*	m["hello"] = 5
		fmt.Println(m["Hello"])

	*/

	//so whats the correct way to add elements
	m = map[string]int{}
	if _, ok := m["Hello"]; ok == false {
		m["Hello"] = 5
	}
	fmt.Println(m["Hello"])

	//as a pratice use these two methods
	var map1 = map[string]int{} //empty map    map[]
	fmt.Println("map1 ", map1)
	map1["Hello"] = 1
	fmt.Println("map1 ", map1)

	map2 := make(map[string]int) //empty map    map[]
	fmt.Println("map2 ", map2)
	map2["Hello"] = 2
	fmt.Println("map2 ", map1)

	//delete from map
	delete(map1, "HOHO") //no error
	fmt.Println("map1 after del ", map1)
	// correct usage
	if _, ok := map1["Hello"]; ok {
		delete(map1, "Hello")
	}
	fmt.Println("map1 after del ", map1)

}
