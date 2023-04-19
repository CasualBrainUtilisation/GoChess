package Moves

import (
	"fmt"
	"strings"

	"github.com/CasualBrainUtilisation/GoChess/Fields"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

// those are the column's letters, in order, they are also used to represents fields e. g.: a6, we'll use this array in the TryToPerformMove function
var columnLetters [8]string = [8]string{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
}

// enums
type MoveType int //represents an enum, that stores the type of a move, which is useful for performing it correctly or to calculate certain things (like in the case of a doublePawnMove the possible en Passant moves)

const (
	Normal         MoveType = 0
	DoublePawnMove MoveType = 1
	EnPassant      MoveType = 2
	Castle         MoveType = 3
)

type Move struct { //a stuct to represent a move on the chessBoard, it is used to store and perform a move
	startPos Fields.BoardField //the position the move starts at
	endPos   Fields.BoardField //the position the move ends at aka the piece moves to

	moveType MoveType //the type of moved performed, e.g.: Normal/Castle ect., needed to perform a move properly or to calculate things like en Passant moves
}

func TryToPerformMove(move string, boardPosition []Pieces.Piece) { //this function will try to perform move from given move notation, if possible

	var curIndex int = 0

	if index := strings.Index(strings.Join(columnLetters[:], ""), strings.ToLower(string(move[curIndex]))); index != -1 { //get the index of the move string as lowercase (as the columnLetters are lowerCase and r n we just check wether the moveString/char represents a column letter) in the columnLetters (converted to a slice, so we can join it to one string) array as joined string, if it does exist (it'll be -1 if it doesn't) continue, adding the x.pos at corresponding position
		fmt.Println(index)
	}
}
