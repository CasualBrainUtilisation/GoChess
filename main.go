package main //this is the main file, meaning it'll handle all the other packages, which theirself will handle most chess game logic

import (
	"github.com/CasualBrainUtilisation/GoChess/Board"
	"github.com/CasualBrainUtilisation/GoChess/FEN"
	"github.com/CasualBrainUtilisation/GoChess/Moves"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

func main() {

	var pieces []Pieces.Piece = FEN.LoadPositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	Board.VisualisePositionFromPieces(pieces)

	Moves.TryToPerformMove("Na6", []Pieces.Piece{})
	Moves.TryToPerformMove("a6", []Pieces.Piece{})
	Moves.TryToPerformMove("6", []Pieces.Piece{})
	Moves.TryToPerformMove("e5", []Pieces.Piece{})
}
