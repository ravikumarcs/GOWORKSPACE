package main

import "fmt"

func binary_search(values []int, value int, low int, high int) int {
	// If not found. Return the negative number of
	// the index where it could possibly go
	if low > high {
		return -low
	}

	mid := low + (high-low)/2

	if values[mid] == value {
		return mid
	} else if values[mid] < value {
		return binary_search(values, value, mid+1, high)
	} else {
		return binary_search(values, value, low, mid-1)
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 17}
	fmt.Println(binary_search(values, 2, 0, len(values)-1))
	fmt.Println(binary_search(values, 7, 0, len(values)-1))
	fmt.Println(binary_search(values, 18, 0, len(values)-1))
}
