package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read file
	dat, err := ioutil.ReadFile("/tmp/abc.txt")
	check(err)

	// split into words
	words := strings.Fields(string(dat))
	// fmt.Println(words)

	// build word distribution
	word_dist := make(map[string]int)
	for _, word := range words {
		//fmt.Println(index , " " , word)

		count, ok := word_dist[word]
		if ok {
			word_dist[word] = count + 1
		} else {
			word_dist[word] = 1
		}
	}

	fmt.Println("Final Word Count")
	for k, v := range word_dist {
		fmt.Printf("[%s] : %d\n", k, v)
	}
}
