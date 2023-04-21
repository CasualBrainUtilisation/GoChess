package Moves

import (
	"github.com/CasualBrainUtilisation/GoChess/Board"
	"github.com/CasualBrainUtilisation/GoChess/Fields"
	"github.com/CasualBrainUtilisation/GoChess/Pieces"
)

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

func getPossibleMovesForPiece(board Board.ChessBoard, piece Pieces.Piece) (moves []Move) { //function that returns all possible moves for given piece on given chessBoard, these moves will not be removed if they are invalid because they cause a check for first, that'll be done in another function

	switch piece.PieceType { //run the corresponding getMoves function for every piece, and add the results to the moves list
	case Pieces.Bishop:
		moves = append(moves, getBishopMoves(board, piece)...)
	case Pieces.Rook:
		moves = append(moves, getRookMoves(board, piece)...)
	case Pieces.Queen:
		moves = append(moves, getQueenMoves(board, piece)...)
	}

	return moves
}

func getBishopMoves(board Board.ChessBoard, piece Pieces.Piece) (moves []Move) { //function that returns all the possible moves for a given piece on given chessBoard, if it was a bishop

	//add the possible moves foreach line a bishop can move on
	moves = append(moves, getLineMoves(board, piece, 1, 1)...)
	moves = append(moves, getLineMoves(board, piece, -1, 1)...)

	return moves //return the calculated moves
}

func getRookMoves(board Board.ChessBoard, piece Pieces.Piece) (moves []Move) { //function that returns all the possible moves for a given piece on given chessBoard, if it was a rook

	//add the possible moves foreach line a rook can move on
	moves = append(moves, getLineMoves(board, piece, 1, 0)...)
	moves = append(moves, getLineMoves(board, piece, 0, 1)...)

	return moves //return the calculated moves
}

func getQueenMoves(board Board.ChessBoard, piece Pieces.Piece) (moves []Move) { //function that returns all the possible moves for a given piece on given chessBoard, if it was a queen

	//the queen can move on the squares a bishop and a rook can move on, so just get the rook and bishop moves and return them
	moves = append(moves, getBishopMoves(board, piece)...)
	moves = append(moves, getRookMoves(board, piece)...)

	return moves //return the calculated moves
}

func getLineMoves(board Board.ChessBoard, piece Pieces.Piece, xIncr, yIncr int) (moves []Move) { //function that returns a list of all possible moves for any given piece on given borad on a line with given gradient, this should be used to get bishop rook and queen moves, notice that this will check the xIncr and yIncr gradient, but also the -xIncr, -yIncr gradient

	moves = append(moves, getMovesForLinePart(board, piece, xIncr, yIncr)...)   //first add the moves possible on line with given gradient
	moves = append(moves, getMovesForLinePart(board, piece, -xIncr, -yIncr)...) //add the moves possible on line in direction opposite to the given gradient

	return moves //return the calculated moves
}

func getMovesForLinePart(board Board.ChessBoard, piece Pieces.Piece, xIncr, yIncr int) (moves []Move) { //method that returns moves for each field on a line with given gradient not the opposite direction though

	var checkedPos Fields.BoardField = piece.BoardPosition //used to check each indivudual field on the line with given gradient for validation
	//already increase checkedPos x and y as we do not want to start checking for line moves on the starting square, which can not be a move
	checkedPos.X += xIncr
	checkedPos.Y += yIncr

	for checkedPos.X >= 0 && checkedPos.X < 8 && checkedPos.Y >= 0 && checkedPos.Y < 8 { //for loop that'll scoute through every possible move destination field on line with given gradient, for that matter of fact, it'll end once the checked pos is out of the board

		moves = append(moves, Move{StartPos: piece.BoardPosition, EndPos: checkedPos, MoveType: Normal})

		//setup the checking for the next field on the line
		checkedPos.X += xIncr
		checkedPos.Y += yIncr
	}

	return moves //return the calculated moves
}

func GetMovesForPieceTypeOfColor(board Board.ChessBoard, pieceType Pieces.PieceType, pieceColor Pieces.PieceColor) (moves []Move) { //function that returns all the moves for all the pieces on given board with given type of given color, this is necessary to get the move a moveNotation is reffering to e. g.: Nf3 --> move with f3 dest, and a Knight moving, it is used in the MoveNotation class for that matter of fact, so it has to be public (capital)

	for _, piece := range board.CurPieces { //foreach piece we'll get the moves and add them to the later returned moves list, if it has the right color and type
		if piece.PieceType == pieceType && piece.PieceColor == pieceColor { //check wether piece is of given type and color
			moves = append(moves, getPossibleMovesForPiece(board, piece)...) //get and add the moves for the piece to the moves list
		}
	}

	return moves //return the calculated moves
}
