package main

import "fmt"

func main() {
	//In Go, nil is not a type but a reserved word. A variable initialized to nil must have a type. I
	//n := nil  //Cannot use 'nil' as the type Type
	var m map[string]int //nil map
	var n *int = nil
	fmt.Println(m, n)
}
