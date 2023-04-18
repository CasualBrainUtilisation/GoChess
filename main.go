package main //this is the main file, meaning it'll handle all the other packages, which theirself will handle most chess game logic

import (
	"fmt"

	"github.com/CasualBrainUtilisation/GoChess/FEN"
	"github.com/CasualBrainUtilisation/GoChess/Fields"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

func main() {
	var piece Pieces.Piece = Pieces.Piece{PieceType: Pieces.Bishop, PieceColor: Pieces.Black, BoardPosition: Fields.BoardField{X: 1, Y: 2}}

	fmt.Println(piece)

	var pieces []Pieces.Piece = FEN.LoadPositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	fmt.Println(pieces)
	fmt.Println("♔", "♚")
}
