package Input //package that manages the user input from the terminal an calls corresponding functions from other packages e. g. = game load "" / Na4 ect.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RespondToUserInputRepeatly() { //loop that'll constantly run, it'll wait for user input, after doing what the user comanded it'll wait for new input ectinput anymore
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
			if len(commandParts) == 0 { //if there is no subcommand, give the user a brief description of the game command
				fmt.Println("game is used to start a new game or load a position from a FEN")
			} else {
				switch strings.ToLower(commandParts[1]) { //search for valid subcommands
				case "new":
					fmt.Println("starting a new game")
				default:
					fmt.Println("invalid subcommand for 'game'")
				}
			}
		}
	}
}
