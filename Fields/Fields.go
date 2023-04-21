package Fields

type BoardField struct { //struct representing any field on the chessBoard, used primarly to represent a piece's position
	X, Y int
}

func IsFieldOnBoard(field BoardField) bool { //checks wether given BoardField even is on a chessBoard (aka x and y are between 8 and 0)
	return field.X >= 0 && field.X < 8 && field.Y >= 0 && field.Y < 8
}
