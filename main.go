package main //this is the main file, meaning it'll handle all the other packages, which theirself will handle most chess game logic

import (
	"fmt"

	"github.com/CasualBrainUtilisation/GoChess/Board"
	"github.com/CasualBrainUtilisation/GoChess/FEN"
	"github.com/CasualBrainUtilisation/GoChess/Input/MoveNotation"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

func main() {

	var pieces []Pieces.Piece = FEN.LoadPositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	Board.VisualisePositionFromPieces(pieces)

	if move, ok := MoveNotation.TryToGetMoveFromNotation("a6", []Pieces.Piece{}); ok == true {
		fmt.Println(move.EndPos)
	}
}
