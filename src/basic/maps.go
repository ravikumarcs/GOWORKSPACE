package main

import "fmt"

func main() {
	num_map := make(map[string]int)

	num_map["one"] = 1
	num_map["two"] = 2

	fmt.Println("map:", num_map)

	fmt.Println("Iterating over map")
	for k, v := range num_map {
		fmt.Printf("key[%s] : value[%d]\n", k, v)
	}
}
