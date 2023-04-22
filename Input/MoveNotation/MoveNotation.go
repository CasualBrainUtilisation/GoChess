package MoveNotation //package that manages user input, if it is a moveNotation for exapmle e5, it'LL get the move the user refers to and perform it with help of the Moves package
import (
	"strconv"
	"strings"
	"unicode"

	"github.com/CasualBrainUtilisation/GoChess/Board"
	"github.com/CasualBrainUtilisation/GoChess/Fields"
	"github.com/CasualBrainUtilisation/GoChess/Moves"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

// those are the column's letters, in order, they are also used to represents fields e. g.: a6, we'll use this array in the TryToGetMoveFromNotation function
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

var pieceNotationLetters [8]string = [8]string{ //array representing the letters sued to specify the piece to move in the SAN move notation in the order of the PieceType enum vales in the Pieces package
	"P", //the pawn can normally not be specified using SAN, further more it'll be used if there is no piece specified, but it doesn't matter, if for some reason you want to, you can do so :)
	"N",
	"B",
	"R",
	"Q",
	"K",
}

func TryToGetMoveFromNotation(board Board.ChessBoard, moveNotation string) (move Moves.Move, ok bool) { //this function will try to get move from given move notation, if it doesn't seem to get a valid move out of it, it'll return false as ok
	ok = false          //set ok to true, so we can naked return, if something doesn't work, later on, we'll change it to true, before the last naked return
	move = Moves.Move{} //this variable will later be returned, as the move represented by the moveNotation, its variables will be calculated throughout this function

	var curIndex int = 0

	//remove any character that might be added in between pieceTypeCharacter and chessBoardFields, that are not necessary to perform moves properly
	moveNotation = strings.ReplaceAll(moveNotation, "-", "")
	moveNotation = strings.ReplaceAll(moveNotation, "x", "")

	var pieceTypeToMove Pieces.PieceType = Pieces.Pawn                                                                  //set the piece type, that will be moved to be a pawn on default, we'll change it if the piece is otherwhise defined in the SAN (it's pawn if there is no specification)
	if index := strings.Index(strings.Join(pieceNotationLetters[:], ""), string(moveNotation[curIndex])); index != -1 { //get the index of the first moveNotation letter in the  joined pieceNotationLetters, if it's -1 (meaning it does not exist in the pieceNotationLetters), just ignore it, else set the pieceTypeToMove to PieceType with the index and increase the curIndex (as now we check the 2nd character)
		pieceTypeToMove = Pieces.PieceType(index) //set the peiceTypeToMove to the one reffered to in the first character fo the SAN
		curIndex += 1                             //increase the curIndex, as we'll now look at the next character for the field Notation
	}

	//calculating the endPos of the move
	if endPos, valid := getFieldPositionFromFieldNotation(moveNotation[curIndex : curIndex+2]); valid == true { //if the notation at the current index (and the next one) is a fieldNotation, set it to the move.EndPos, else return (as the notation is invalid, cuz there has to be a fieldNotation here)
		move.EndPos = endPos
	} else {
		return
	}

	if movesForPieceTypeNoted := Moves.GetMovesForPieceTypeOfColor(board, pieceTypeToMove, Pieces.White); len(movesForPieceTypeNoted) != 0 { //get the moves for the pieceType the moveNotation is reffering to and check wether, there are any
		for _, moveForPieceTypeNoted := range movesForPieceTypeNoted { //foreach move the pieceType the moveNotation is reffering to can do, check wether it has the in the moveNotation given endPos
			if moveForPieceTypeNoted.EndPos == move.EndPos { //if there is a move on the board with given endPos, we'll set the move's start pos to its startPos, also set ok to true as the move is valid, it'll be false on default, so if there is no move with the endPos we'll later return false
				move.StartPos = moveForPieceTypeNoted.StartPos //set the move to use the found move for piece type as start pos
				ok = true                                      //set ok to true as the move is valid now
			}
		}
	}

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
