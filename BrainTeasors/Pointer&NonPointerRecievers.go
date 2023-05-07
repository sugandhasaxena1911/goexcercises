package main

import "fmt"

type employee struct {
	fname string
	id    int
}

type exampletype struct {
	n int
}

type intf interface {
	M()
}

func main() {
	e := employee{
		fname: "Sugandha",
		id:    123,
	}
	fmt.Println("Using non pointer to call both POINTERS & NON POINTER methods")
	fmt.Println("---------------------------------------------------------------")
	e.A()
	fmt.Println(e)

	e.B()
	fmt.Println(e)

	fmt.Println("Using  pointer to call both POINTERS & NON POINTER methods")
	fmt.Println("---------------------------------------------------------------")
	e1 := employee{
		fname: "Suggu",
		id:    125,
	}
	p := &e1
	p.A()
	fmt.Println(e1)

	p.B()
	fmt.Println(e1)

	fmt.Println("Functions M belongs to interface ")
	fmt.Println("-------------------------------------")
	e3 := employee{
		fname: "Saxena",
		id:    121,
	}
	fmt.Println("m is defined by employee")
	fmt.Println("---------------------------------------------------------------")
	e3.M()
	fmt.Println(e3)
	r := &e3
	r.M()
	fmt.Println(e3)
	// employee and pointer of employee both impeemntts interface i
	var i intf
	i = e3
	fmt.Println(i)

	i = r
	fmt.Println(i)

	et1 := exampletype{1}
	et2 := exampletype{2}
	rt := &et2

	et1.M()
	fmt.Println(et1.n)
	rt.M()
	fmt.Println(et2.n)
	i = et1 //et1 doesnt implement interface.Only pointer of type impleents interface
	i = rt

}

func (emp employee) A() {
	emp.fname = "Changed HUHUHU name"
	fmt.Println("I am in A - Non pointer reciever ", emp.fname, emp.id)
}
func (emp *employee) B() {
	emp.fname = "Changed name"
	fmt.Println("I am in B -  pointer reciever", emp.fname, emp.id)
}

func (emp employee) M() {
	emp.fname = "HAHAHHA"
	fmt.Println("I am in M - Non pointer reciever ", emp.fname, emp.id)

}
func (et *exampletype) M() {
	fmt.Println(" HAHHAA i in M ", et)
	et.n = 5
}
