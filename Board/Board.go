package Board //package to manage the chessBoard, e. g. : the piece characters

import (
	"fmt"

	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

// Now following all the chessPiece characters, again sorted so the Piece.PieceType can be used as index, those characters are official Unicode characters
var whiteUnicodeChessPieces [6]string = [6]string{
	"♙",
	"♘",
	"♗",
	"♖",
	"♕",
	"♔",
}

func VisualisePositionFromPieces(position []Pieces.Piece) { //shows the current position on the chessBoard in the terminal utilising the chessPiece unicoode characters

	var visualisation [8][8]string //create the array that'll later be printed out to visualize the board

	for y := 0; y < 8; y++ { //add the empty fields, to the visualisation chessBoard array
		for x := 0; x < 8; x++ {
			visualisation[y][x] = "_"
		}
	}

	//ovverride certain empty fields in the visualisation array with the corresponding piece character
	for _, piece := range position {
		visualisation[piece.BoardPosition.Y][piece.BoardPosition.X] = whiteUnicodeChessPieces[piece.PieceType]
	}

	//print out the board, each row will be printed in a new line
	for _, row := range visualisation {
		fmt.Println(row)
	}
}
