package FEN //File that manages loading positions out of a FEN-string, or saving positions to a FEN-string

import (
	"strconv"
	"strings"
	"unicode"

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

		if column > 8 || row > 8 { //check wether the current row or column is over 8, notice that values of 8 are valid, as they might be followed by a slash (letting 8 go through adding a piece to the toReturn slices is ok, as the new column value of 9 will immeadatly cause this if statement to be true anyhow) (over 8 is off board for sure though), if row/column are over 8, return an empty Pieces.Piece slice
			return make([]Pieces.Piece, 0, 0)
		}

		var curString string = string(fenParts[0][i])                                                           //get the current string, so we do not have to get it every time we need it
		if index := strings.Index(strings.Join(pieceLetters[:], ""), strings.ToLower(curString)); index != -1 { //get the index of the current string as lowercase (as the pieceLetters are lowerCase and r n we just check wether the curString/char represents a chess piece) in the pieceLetters (converted to a slice, so we can join it to one string) array as joined string, if it does exist (it'll be -1 if it doesn't) continue, adding the piece at corresponding position to the toReturn slice
			//calculate the index for the PieceColor enum, of the piece we are about to create
			var colorInd int = 1                         //first set the piece to black, it'll be set to white if needed
			if strings.ToUpper(curString) == curString { //check if the curString is uppercase, by checking wether 'upering' the curString is still giving the same string as curString
				colorInd = 0 //if the curString is uppercase, we have to set the colorIndex to 0, so the piece will be white
			}

			toReturn = append(toReturn, Pieces.Piece{PieceType: Pieces.PieceType(index), PieceColor: Pieces.PieceColor(colorInd), BoardPosition: Fields.BoardField{X: column, Y: row}}) //add the piece from curString, set its type to the index of the curString in the pieceLetters array, set the color to the calculated colorInd(ex) and set the position to the current column and row, of the FEN string

			column += 1 //increase the column by 1, so the next piece is placed in the next column
			continue    //continue with the next character of the FEN string
		} else if unicode.IsNumber(rune(curString[0])) { //check wether the character of curString, is a number, if so increase the column by it
			columnIncr, err := strconv.Atoi(curString) //convert the string to a number
			if err != nil {                            //if for some odd reason the string could not be converted to a number (which I could not understand whatshowever) return an empty Piece.Piece slice, as the FEN seems invalid
				return make([]Pieces.Piece, 0, 0)
			}
			column += columnIncr //increase the column by the curString represented int
		} else if curString == "/" { //if the curString is a slash, we'll increase the row by 1 and reset the column, to fill the next board row
			row += 1   //increase the row
			column = 0 //reset the column so we start at column 0 in the new row
		} else { //if the currentString is not a valid FEN character we return an empty Piece.Piece slice, as the FEN
			return make([]Pieces.Piece, 0, 0)
		}
	}

	return toReturn //return created list
}
