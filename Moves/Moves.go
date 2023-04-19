package Moves

import (
	"fmt"
	"strings"

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

func TryToPerformMove(move string, boardPosition []Pieces.Piece) {

	var curIndex int = 0

	if index := strings.Index(strings.Join(columnLetters[:], ""), strings.ToLower(string(move[curIndex]))); index != -1 { //get the index of the move string as lowercase (as the columnLetters are lowerCase and r n we just check wether the moveString/char represents a column letter) in the columnLetters (converted to a slice, so we can join it to one string) array as joined string, if it does exist (it'll be -1 if it doesn't) continue, adding the x.pos at corresponding position
		fmt.Println(index)
	}
}
