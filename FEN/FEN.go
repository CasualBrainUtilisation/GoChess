package FEN //File that manages loading positions out of a FEN-string, or saving positions to a FEN-string
import "github.com/CasualBrainUtilisation/GoChess/Pieces"

func LoadPositionFromFEN(fen string) []Pieces.Piece {

	var toReturn []Pieces.Piece = make([]Pieces.Piece, 0, 0) //create list with ChessPieces, representing a chess position, that'll later be returned

	return toReturn //return created list
}
