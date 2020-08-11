package main

import "fmt"

func main() {
	board := CreateFromFEN("0/0/0/0/0/0/0/0")
	for i := 0; i < len(board.pieces); i++ {
		if i%8 == 0 && i != 0 {
			fmt.Print("\n")
		}

		if board.pieces[i] == 0 {
			fmt.Print("_" + " ")
		} else {
			fmt.Print(string(board.pieces[i]) + " ")
		}
	}
}
