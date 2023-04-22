package Input //package that manages the user input from the terminal an calls corresponding functions from other packages e. g. = game load "" / Na4 ect.
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CasualBrainUtilisation/GoChess/Board"
	"github.com/CasualBrainUtilisation/GoChess/FEN"
)

func RespondToUserInputRepeatly(chessBoard Board.ChessBoard) { //loop that'll constantly run, it'll wait for user input, after doing what the user comanded it'll wait for new input ect, it needs the chessBoard for a bunch of function it'll run as response to some inputs
	reader := bufio.NewReader(os.Stdin) //reader that'll be used throughout this function, to get the user input in the terminal

	for { //while this programm is supposed to run (the user will stop that eventually), we'll get the user's input with the help of the bufio.Reader and run the corresponding functions here
		input, _ := reader.ReadString('\n')     //wait til the user inputs something in the terminal with a newline (\n), after that we'll get the input
		input = strings.TrimSuffix(input, "\n") //remove the newLine at the end from the input, so only the actual input is left

		var commandParts []string = strings.Fields(input) //get the parts of the user's input, divided by whitespaces, so we can access subcommands easily

		switch strings.ToLower(commandParts[0]) { //check if the user inputted certain things and if so run the corresponding functions (compare lowercased so the case doesn't matter)
		case "quit", "stop", "exit": //in case the user put in 'quit' ect, we'll end this programm by ending the function with return
			fmt.Println()
			fmt.Println("Ending the game")
			return
		case "game": //game is used to start new games and load FENs ect. for that we'll search for further subcommands
			if len(commandParts) == 1 { //if there is no subcommand, give the user a brief description of the game command
				fmt.Println("game is used to start a new game or load a position from a FEN")
			} else {
				switch strings.ToLower(commandParts[1]) { //search for valid subcommands
				case "new":
					fmt.Println("starting a new game") //inform the user that its command'll be executed
					fmt.Println()

					chessBoard = Board.ChessBoard{CurPieces: FEN.LoadPositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")} //create a chessBoard with the start position loaded from the start FEN
					chessBoard.VisualisePositionFromPieces()                                                                                      //print out chessBoard
				case "load": //this will load the FEN in the commandPart[2]
					if len(commandParts) == 2 { //if there is no subcommand to load, give the user a brief description of the game load command
						fmt.Println("load can be used to load a FEN on the CLI chessBoard")
					} else {
						chessBoard = Board.ChessBoard{CurPieces: FEN.LoadPositionFromFEN(commandParts[2])} //load the FEN on a new chessBoard, we'll store in the chessBoard variable
						chessBoard.VisualisePositionFromPieces()                                           //print out loaded chessBoard
					}
				default:
					fmt.Println("invalid subcommand for 'game'")
				}
			}
		}

		fmt.Println() //always print a line here to seperate the user input line from the new output
	}
}
