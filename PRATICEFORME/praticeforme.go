package PRATICEFORME

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func BinarySearch1(datalist []int, data int) int {

	low := 0
	high := len(datalist) - 1

	for low <= high {
		mid := high + low/2
		if datalist[mid] == data {
			return mid
		}
		if data > datalist[mid] {
			low = mid + 1
		}
		if data < datalist[mid] {
			high = mid - 1
		}

	}
	return -1
}

func Bubble1(datalist []int) []int {
	n := len(datalist) - 1
	sorted := false
	for !sorted {
		swapped := false
		for x := 0; x < n; x++ {
			if datalist[x] > datalist[x+1] {
				datalist[x], datalist[x+1] = datalist[x+1], datalist[x]
				swapped = true
			}

		}
		if !swapped {
			sorted = true

		}
		n = n - 1

	}
	return datalist
}

func QuickPartitions(datalist []int, low int, high int) ([]int, int) {
	pivot := datalist[high]
	pointer := low

	for x := low; x < high; x++ {

		if datalist[x] < pivot {
			//swap
			datalist[x], datalist[pointer] = datalist[pointer], datalist[x]
			pointer++
		}
	}

	datalist[pointer], datalist[high] = datalist[high], datalist[pointer]
	return datalist, pointer

}

func QuickSort1(datalist []int, low, high int) []int {
	if low < high {
		datalist, pointer := QuickPartitions(datalist, low, high)
		QuickSort1(datalist, low, pointer-1)
		QuickSort1(datalist, pointer+1, high)

	}
	return datalist
}

func MoveZeroes(data []int) []int {
	pointer := 0
	for x := 0; x < len(data); x++ {
		if data[x] != 0 {
			data[pointer], data[x] = data[x], data[pointer]
			pointer++
		}
	}
	return data
}

func MergeList(data [][]int) [][]int {

	final := [][]int{}
	prev := data[0]
	for x := 1; x < len(data); x++ {
		current := data[x]
		if current[0] == prev[1] {
			// create merged node
			node := []int{prev[0], current[1]}
			prev = node

		} else {
			//append current node into list
			final = append(final, prev)
			prev = current
		}

	}
	final = append(final, prev)
	return final
}

func HandleRequests() {
	rt := mux.NewRouter()
	rt.HandleFunc("/home", HomePage).Methods("POST")
	log.Fatalln(http.ListenAndServe(":5500", rt))

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Homepage"))
}
