package main

// store a chessboard position
type ChessBoard struct {
	pieces [64]byte
}

// create a fen string from a chess position
func CreateFromFEN(fen string) ChessBoard {
	return ChessBoard{pieces: [64]byte{}}
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
			str += string(board.pieces[i])
		}
	}
	return str
}
