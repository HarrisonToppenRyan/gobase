package main

// store a chessboard position
type ChessBoard struct {
	pieces [64]byte
}

// create a fen string from a chess position
func CreateFromFEN(fen string) (board ChessBoard) {
	board = ChessBoard{pieces: [64]byte{}}

	var position int = 0
	for i := 0; i < len(fen); i++ {
		if fen[i] == '/' {
			continue
		}

		// check if fen character is a number
		if fen[i] < 57 && fen[i] > 48 {
			// get blank square number
			num := int(fen[i]) - 48

			// add blank squares
			for j := 0; j < num; j++ {
				board.pieces[position] = 0
				position++
			}
		} else {
			board.pieces[position] = fen[i]
			position++
		}
	}

	return board
}

// get the string form of a chessboard
func String(board ChessBoard) (str string) {
	for i := 0; i < len(board.pieces); i++ {
		if i%8 == 0 && i != 0 {
			str += "\n"
		}

		if board.pieces[i] == 0 {
			str += "_ "
		} else {
			str += string(string(board.pieces[i]) + " ")
		}
	}
	return str
}
