package main

import "fmt"

func main() {
	word := "selamat malam"
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}
