package main

import "fmt"

func main() {
	word := "selamat malam"
	calculate := make(map[string]int)
	for i := 0; i < len(word); i++ {
		letter := string(word[i])
		calculate[letter]++
		fmt.Println(letter)
	}
	fmt.Println(calculate)
}
