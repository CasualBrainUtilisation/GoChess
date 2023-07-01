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

type ChessBoard struct { //struct that represents a chessBoard, therefore storing the piece's currently on board
	CurPieces   []Pieces.Piece    //variable that stores the pieces currently placed on the chess board
	ColorToMove Pieces.PieceColor //variable that stores the color that will perform the next move
}

func (chessBoard ChessBoard) VisualisePositionFromPieces() { //shows the current position on the chessBoard in the terminal utilising the chessPiece unicoode characters

	var visualisation [8][9]string //create the array that'll later be printed out to visualize the board, notice that this has 9 characters (not 8), wo we have an empty space after the pieces

	for y := 0; y < 8; y++ { //add the empty fields, to the visualisation chessBoard array
		for x := 0; x < 8; x++ {
			visualisation[y][x] = " "
		}
	}

	//ovverride certain empty fields in the visualisation array with the corresponding piece character
	for _, piece := range chessBoard.CurPieces {
		var pieceCharacter string = blackUnicodeChessPieces[piece.PieceType] //set the piece character, that'll use to visualize the piece on the chessBoard, r n we set it to the black piece character of corresponding type
		if piece.PieceColor == Pieces.White {                                //check if the piece is colored in white, if so set the pieceCharacter to the white counterPart
			pieceCharacter = whiteUnicodeChessPieces[piece.PieceType]
		}

		visualisation[piece.BoardPosition.Y][piece.BoardPosition.X] = pieceCharacter //ovveride the character in the visualized board grid at the piece position with the evaluated pieceCharacter
	}

	//print out the board, each row will be printed in a new line
	for i, row := range visualisation {
		fmt.Println(8-i, " ", row) //8-i prints the row number while 'row' prints the pieces in the row, also we add a space between those for clarity
	}
	fmt.Println()
	fmt.Println("   ", [9]string{"A", "B", "C", "D", "E", "F", "G", "H", ""}) //print out the letters to indicate the column position of pieces
}

func (chessBoard ChessBoard) GetPieceAtBoardPosition(x, y int) (pieceAtPos *Pieces.Piece, ok bool) { //function that'll return the piece at given x,y position on the chessBoard, ok'll be false, if there is no piece at given pos
	ok = false                                            //set ok to false on default so it'll be false if there is no piece found at given x,y position
	for pieceIndex, piece := range chessBoard.CurPieces { //check foreach piece wether it is placed at the given x, y position
		if piece.BoardPosition.X == x && piece.BoardPosition.Y == y { //check wether the piece is positioned at given board Position
			pieceAtPos = &chessBoard.CurPieces[pieceIndex] //set the pieceAtPos (will be returned) to piece found at given x, y position, use the pointer, so that other functions can use this method and modify returned piece, this for example is needed to move pieces
			ok = true                                      //set ok to true as we found a piece
		}
	}
	return //return the results
}

func (chessBoard ChessBoard) RemovePieceFromBoard(pieceToRemove *Pieces.Piece) { //functino that'll remove the piece from the board.CurPieces list, that is placed at given position
	for i, piece := range chessBoard.CurPieces {
		if piece == *pieceToRemove {
			chessBoard.CurPieces = append(chessBoard.CurPieces[:i], chessBoard.CurPieces[i+1:]...)
		}
	}
}
