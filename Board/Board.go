package Board //package to manage the chessBoard, e. g. : the piece characters

import (
	"fmt"

	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

// Now following all the chessPiece characters, again sorted so the Piece.PieceType can be used as index, those characters are official Unicode characters
// black chess symbols
var blackUnicodeChessPieces [6]string = [6]string{
	"♙",
	"♘",
	"♗",
	"♖",
	"♕",
	"♔",
}

// white chess symbols
var whiteUnicodeChessPieces [6]string = [6]string{
	"♟︎",
	"♞",
	"♝",
	"♜",
	"♛",
	"♚",
}

func VisualisePositionFromPieces(position []Pieces.Piece) { //shows the current position on the chessBoard in the terminal utilising the chessPiece unicoode characters

	var visualisation [8][8]string //create the array that'll later be printed out to visualize the board

	for y := 0; y < 8; y++ { //add the empty fields, to the visualisation chessBoard array
		for x := 0; x < 8; x++ {
			visualisation[y][x] = " "
		}
	}

	//ovverride certain empty fields in the visualisation array with the corresponding piece character
	for _, piece := range position {
		var pieceCharacter string = blackUnicodeChessPieces[piece.PieceType] //set the piece character, that'll use to visualize the piece on the chessBoard, r n we set it to the black piece character of corresponding type
		if piece.PieceColor == Pieces.White {                                //check if the piece is colored in white, if so set the pieceCharacter to the white counterPart
			pieceCharacter = whiteUnicodeChessPieces[piece.PieceType]
		}

		visualisation[piece.BoardPosition.Y][piece.BoardPosition.X] = pieceCharacter //ovveride the character in the visualized board grid at the piece position with the evaluated pieceCharacter
	}

	//print out the board, each row will be printed in a new line
	for _, row := range visualisation {
		fmt.Println(row)
	}
}
