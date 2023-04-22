package main //this is the main file, meaning it'll handle all the other packages, which theirself will handle most chess game logic

import (
	"fmt"

	"github.com/CasualBrainUtilisation/GoChess/Board"
	"github.com/CasualBrainUtilisation/GoChess/FEN"
	"github.com/CasualBrainUtilisation/GoChess/Input"
	"github.com/CasualBrainUtilisation/GoChess/Input/MoveNotation"
)

func main() {

	var chessBoard = Board.ChessBoard{CurPieces: FEN.LoadPositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")} //create a chessBoard with the start position loaded from the start FEN

	chessBoard.VisualisePositionFromPieces()

	if move, ok := MoveNotation.TryToGetMoveFromNotation(chessBoard, "d3"); ok == true {
		fmt.Println(move.EndPos)
	}

	Input.RespondToUserInputRepeatly()
}
