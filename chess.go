package main

// store a chessboard position
type ChessBoard struct {
	pieces [64]byte
}

// create a fen string from a chess position
func CreateFromFEN(fen string) ChessBoard {
	return ChessBoard{pieces: [64]byte{}}
}
