package Moves

import (
	"strconv"
	"strings"
	"unicode"

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

const ( //defines all MoveTypes to be used as enum values
	Normal         MoveType = 0
	DoublePawnMove MoveType = 1
	EnPassant      MoveType = 2
	Castle         MoveType = 3
)

type Move struct { //a stuct to represent a move on the chessBoard, it is used to store and perform a move
	StartPos Fields.BoardField //the position the move starts at
	EndPos   Fields.BoardField //the position the move ends at aka the piece moves to

	MoveType MoveType //the type of moved performed, e.g.: Normal/Castle ect., needed to perform a move properly or to calculate things like en Passant moves
}

func TryToGetMove(moveNotation string, boardPosition []Pieces.Piece) (move Move, ok bool) { //this function will try to get move from given move notation, if it doesn't seem to get a valid move out of it, it'll return false as ok
	ok = true     //set ok to true on default, will be set to false if the moveNotation isn't valid throughout this function
	move = Move{} //this variable will later be returned, as the move represented by the moveNotation, its variables will be calculated throughout this function

	var curIndex int = 0

	//calculating the endPos of the move

	if index := strings.Index(strings.Join(columnLetters[:], ""), strings.ToLower(string(moveNotation[curIndex]))); index != -1 { //get the index of the move string as lowercase (as the columnLetters are lowerCase and r n we just check wether the moveString/char represents a column letter) in the columnLetters (converted to a slice, so we can join it to one string) array as joined string, if it does exist (it'll be -1 if it doesn't) continue, adding the x.pos at corresponding position
		move.EndPos.X = index //set the x position of the move endPos to the index of the first letter in the letters, as that is what it represents
	}
	if unicode.IsNumber(rune(moveNotation[1])) { //check wether the 2nd character of moveNotation, is a number, if so set the move.endPos.Y to it, as that is what it represents
		endY, err := strconv.Atoi(string(moveNotation[1])) //convert the string to a number
		if err != nil {                                    //if for some odd reason the string could not be converted to a number (which I could not understand whatshowever) return set ok to false
			ok = false
		} else { //set the move.endPos.Y accordingly if the string could be converted
			move.EndPos.Y = endY
		}

	}

	return //naked return statement, cuz i'm lazy, returns move and ok
}
