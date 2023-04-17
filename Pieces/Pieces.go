package Pieces

import "github.com/CasualBrainUtilisation/GoChess/Fields"

//enums (or the og version of them)
//enum that represents the chessPiece.Type
type PieceType int //type used to represent the type of a chessPiece

const ( //defines all PieceTypes to be used as enum values
	Pawn   PieceType = 0
	Knight PieceType = 1
	Bishop PieceType = 2
	Rock   PieceType = 3
	Queen  PieceType = 4
	King   PieceType = 5
)

//enum that represents the chessPiece.color
type PieceColor int //type used to represent the color of a chessPiece

const ( //defines all PieceColors to be used as enum values
	White PieceColor = 0
	Black PieceColor = 1
)

type Piece struct { //struct that represents any chessPiece from pawn to king
	PieceType  PieceType  //variable representing the type of the piece for example: 1 for Knight, those are defined in the consts() above
	PieceColor PieceColor //variable representing the color of the piece for example: 0 for Black, those are defined in the consts() above

	BoardPosition Fields.BoardField //variable representing the position of the chessPiece on the chessBoard
}
