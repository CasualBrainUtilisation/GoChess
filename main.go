package main //this is the main file, meaning it'll handle all the other packages, which theirself will handle most chess game logic

import (
	"github.com/CasualBrainUtilisation/GoChess/Fields"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

func main() {
	var piece Pieces.Piece = Pieces.Piece{Pieces.Bishop, Pieces.Black, Fields.BoardField{1, 2}}
}
