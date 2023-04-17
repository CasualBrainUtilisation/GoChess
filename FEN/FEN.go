package FEN //File that manages loading positions out of a FEN-string, or saving positions to a FEN-string

import (
	"strings"

	"github.com/CasualBrainUtilisation/GoChess/Fields"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

// arrays
var pieceLetters [6]string = [6]string{ //array representing all letters, which represent a chessPiece in the FEN, the strings are orders, so they can be used as an index for the PieceType enum in the Pieces package
	"p",
	"n",
	"b",
	"r",
	"q",
	"k",
}

func LoadPositionFromFEN(fen string) []Pieces.Piece {

	var toReturn []Pieces.Piece = make([]Pieces.Piece, 0, 0) //create list with ChessPieces, representing a chess position, that'll later be returned

	var fenParts []string = strings.Fields(fen) //split the FEN string into umltiple pieces, all determinated by the spaces

	//check the first piece of the splitted FEN string, which represents the pieces and their position on the board
	for i, column, row := 0, 0, 0; i < len(fenParts[0]); i++ { //we'll check each character of the first fenPart piece, also we'll create the column and row variables to use in the for loop, to store the current position of the currently checked character's piece
		var curString string = string(fenParts[0][i])                                                           //get the current string, so we do not have to get it every time we need it
		if index := strings.Index(strings.Join(pieceLetters[:], ""), strings.ToLower(curString)); index != -1 { //get the index of the current string as lowercase (as the pieceLetters are lowerCase and r n we just check wether the curString/char represents a chess piece) in the pieceLetters (converted to a slice, so we can join it to one string) array as joined string, if it does exist (it'll be -1 if it doesn't) continue, adding the piece at corresponding position to the toReturn slice
			//calculate the index for the PieceColor enum, of the piece we are about to create
			var colorInd int = 1                         //first set the piece to black, it'll be set to white if needed
			if strings.ToLower(curString) == curString { //check if the curString is lowercase, by checking wether lowering the curString is still giving the same string as curString
				colorInd = 0 //if the curString is lowercase, we have to set the colorIndex to 0, so the piece will be white
			}

			toReturn = append(toReturn, Pieces.Piece{PieceType: Pieces.PieceType(index), PieceColor: Pieces.PieceColor(colorInd), BoardPosition: Fields.BoardField{X: column, Y: row}}) //add the piece from curString, set its type to the index of the curString in the pieceLetters array, set the color to the calculated colorInd(ex) and set the position to the current column and row, of the FEN string

			column += 1 //increase the column by 1, so the next piece is placed in the next column
			continue    //continue with the next character of the FEN string
		}
	}

	return toReturn //return created list
}
