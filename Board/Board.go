package Board //package to manage the chessBoard, e. g. : the piece characters

import (
	"fmt"

	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

func VisualisePositionFromPieces(position []Pieces.Piece) { //shows the current position on the chessBoard in the terminal utilising the chessPiece unicoode characters

	var visualisation [8][8]string = [8][8]string{{" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}, {" ", " ", " ", " ", " ", " ", " ", " "}}

	for _, piece := range position {
		visualisation[piece.BoardPosition.Y][piece.BoardPosition.X] = "P"
	}

	for _, row := range visualisation {
		fmt.Println(row)
	}
}
