package main

import "fmt"

func main() {
	board := CreateFromFEN("r1b1k2r/pp3ppp/2pp4/8/2P5/3NPnP1/PPK2P1P/R4B1R")
	fmt.Println(String(board))
}
