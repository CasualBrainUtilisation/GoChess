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
	ok = false    //set ok to true, so we can naked return, if something doesn't work, later on, we'll change it to true, before the last naked return
	move = Move{} //this variable will later be returned, as the move represented by the moveNotation, its variables will be calculated throughout this function

	var curIndex int = 0

	//calculating the endPos of the move
	if endPos, valid := getFieldPositionFromFieldNotation(moveNotation[curIndex:2]); valid == true { //if the notation at the current index (and the next one) is a fieldNotation, set it to the move.EndPos, else return (as the notation is invalid, cuz there has to be a fieldNotation here)
		move.EndPos = endPos
	} else {
		return
	}

	ok = true
	return //naked return statement, cuz i'm lazy, returns move and ok
}

func getFieldPositionFromFieldNotation(fieldNotation string) (fieldPos Fields.BoardField, ok bool) { //function that'll get the position on the board from a fieldNotation, e.g.: a4 --> BoardPos{X=0, Y=4}, it also returns ok, which is false if the fieldNotation was not valid
	fieldPos = Fields.BoardField{} //set the later retunred fieldPos to a new BoardField, which values we'll chagne throughout this function
	ok = false                     //set the later returned ok to false for first, so we can always do a naked return, which will just return ok as false, and we do not have to set it every time, just set it to true before the last return

	if index := strings.Index(strings.Join(columnLetters[:], ""), strings.ToLower(string(fieldNotation[0]))); index != -1 { //get the index of the fieldnotation string as lowercase (as the columnLetters are lowerCase and r n we just check wether the moveString/char represents a column letter (doesn't matter wether cap)) in the columnLetters (converted to a slice, so we can join it to one string) array as joined string, if it does exist (it'll be -1 if it doesn't) continue, adding the x.pos at corresponding position
		fieldPos.X = index //set the x position of the fieldPos to the index of the first letter in the letters, as that is what it represents
	} else { //there has to be a letter here, or the fieldNotation is invalid, so return (ok already false)
		return //return ok=false
	}
	if unicode.IsNumber(rune(fieldNotation[1])) { //check wether the 2nd character of moveNotation, is a number, if so set the fieldPos.Y to it, as that is what it represents
		endY, err := strconv.Atoi(string(fieldNotation[1])) //convert the string to a number
		if err != nil {                                     //if for some odd reason the string could not be converted to a number (which I could not understand whatshowever) return ok=false
			return
		} else { //set the fieldPos.Y accordingly if the string could be converted
			fieldPos.Y = 8 - endY //row 8 is 0y and row 1 is 7y ect. so thats why it's done this way
		}

	}

	ok = true //there hasn't been an issue so ok should be true now
	return    //return the calculated results
}