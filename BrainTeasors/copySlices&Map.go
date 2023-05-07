package main

import "fmt"

func main() {

	sl1 := []int{1, 2, 3, 4}
	fmt.Println("slice address sl1 :", &sl1)
	sl2 := sl1
	fmt.Println("slice address sl2 :", &sl2, sl2)
	var sl3 []int
	sl3 = sl1
	fmt.Println("slice address sl3 :", &sl3, sl3)
	sl2[1] = 999
	sl3[2] = 888
	fmt.Println(sl1, sl2, sl3) // chnages are reflected to all
	changeSlice(sl1)           // slices are pass by reference
	fmt.Println(sl1, sl2, sl3) // chnages are reflected to all
	// REASON : because the underlying array is same
	//copy using builin copy
	sl4 := make([]int, len(sl1))
	copy(sl4, sl1)
	sl4[0] = 1111
	fmt.Println(sl1, sl4)

	//copied upto smaller elements
	sl5 := make([]int, len(sl1)-2)
	copy(sl5, sl1)
	sl5[0] = 1111
	fmt.Println(sl1, sl5)

	//copy to empty s6 or nil s7 slice   : nothing happens
	// With copy, the destination slice will not be nil even if the source slice is nil;:
	//					REASON : basically destination itself here is non nil & it reamains same after that
	sl6 := []int{} //empty slice
	fmt.Println(sl6)

	copy(sl6, sl1)
	//sl6[0] = 2222    //this gives error for index out of bound
	fmt.Println(sl1, sl6)

	var s7 []int
	fmt.Println(s7)
	copy(s7, sl1)
	//s7[0] = 2222 //this gives error for index out of bound
	fmt.Println(sl1, s7)

	// dest remains same as nil or non nil
	var srcsl []int //nil source slice
	destsl := make([]int, len(srcsl))
	copy(destsl, srcsl)
	fmt.Println("srcsl==nil : ", srcsl == nil, "destsl==nil : ", destsl == nil)

	//using append to copy to empty or nil slice
	sl1 = []int{1, 2, 3, 4, 5}
	var sl8 []int
	sl8 = append(sl8, sl1...)
	sl8[0] = 888
	fmt.Println(sl8, sl1)

	srcsl1 := make([]int, 0)
	destsl1 := make([]int, 0)
	destsl1 = append(destsl1, srcsl1...)
	fmt.Println("srcsl==nil : ", srcsl1 == nil, "destsl==nil : ", destsl1 == nil)

}
func changeSlice(sl []int) {

	sl[3] = 555
}
