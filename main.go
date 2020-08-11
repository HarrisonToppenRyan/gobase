package main

import "fmt"

func main() {
	board := CreateFromFEN("0/0/0/0/0/0/0/0")
	fmt.Println(String(board))
}
