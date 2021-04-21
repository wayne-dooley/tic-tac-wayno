/* Tic-Tac-Wayno
*
*  My very first Go language program
*
 */

package main

import (
	"fmt"
)

// Displays the playing board
func showBoard(sb [10]string) {
	fmt.Println("\n  ", sb[1], "|", sb[2], "|", sb[3])
	fmt.Println("  -----------")
	fmt.Println("  ", sb[4], "|", sb[5], "|", sb[6])
	fmt.Println("  -----------")
	fmt.Println("  ", sb[7], "|", sb[8], "|", sb[9])
	fmt.Println(" ")

}

// Display into
func main() {

	// Intro screen
	//
	fmt.Println("\n***********************************")
	fmt.Println("**                               **")
	fmt.Println("**  Welcome to tic-tac-wayno!!!  **")
	fmt.Println("**                               **")
	fmt.Println("***********************************")
	fmt.Println(" ")

	/* Declare the board
	*
	*  Set the board as a 3x3 grid using a 9-position array.
	*  Map each square 1-9:
	*    1 | 2 | 3
	*    4 | 5 | 6
	*    7 | 8 | 9
	*
	*  Set values to track the moves and display
	*    " " = available move
	*    "X" = user move (X)
	*    "O" = computer move (O)
	 */

	// Initialize the board
	board := [...]string{"", " ", " ", " ", " ", " ", " ", " ", " ", " "}

	// Refresh the display
	showBoard(board)

	var xMove int
	var oTurn bool
	var tie bool
	var xWin bool

	var corner = [4]int{1, 3, 7, 9} // Define corners
	var side = [4]int{2, 4, 6, 8}   // Define sides

	// Declare all possible board winning positions
	var b = [8][3]int{
		// diagonal
		{1, 5, 9},
		{3, 5, 7},
		// horizontal
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		// vertical
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9}}

	// Start the game loop
	for {

		/*
		**
		**  User move (X)
		**
		 */

		// Accept user move (X) and validate
		fmt.Print("Player X move : ")
		fmt.Scanf("%d ", &xMove)

		for {
			// Must be 1-9, and square must be available
			if xMove < 1 || xMove > 9 {
				fmt.Print("\n** Must be between 1 and 9, try again : ")
				fmt.Scanf("%d ", &xMove)
			} else if board[xMove] != " " {
				fmt.Print("\n** Move is already taken, try again : ")
				fmt.Scanf("%d ", &xMove)
			} else {
				break
			}
		}

		// Make X move
		board[xMove] = "X"

		// Did X win?
		xWin = false
		for i := 0; i < 8; i++ {
			if board[b[i][0]] == "X" && board[b[i][1]] == "X" && board[b[i][2]] == "X" {
				xWin = true
				break
			}
		}

		if xWin {
			showBoard(board)
			fmt.Println("***************")
			fmt.Println("**           **")
			fmt.Println("**  X WINS!  **")
			fmt.Println("**           **")
			fmt.Println("***************")
			fmt.Println(" ")
			break
		}

		// Tie game?
		tie = true
		for i := 1; i <= 9; i++ {
			if board[i] == " " {
				tie = false
				break
			}
		}

		if tie {
			showBoard(board)
			fmt.Println("****************")
			fmt.Println("**            **")
			fmt.Println("**  Tie Game  **")
			fmt.Println("**            **")
			fmt.Println("****************")
			fmt.Println(" ")
			break
		}

		/*
		**
		**  Computer move (O)
		**
		 */

		oTurn = true

		// Can O make a winning move?
		if oTurn {
			for i := 0; i < 8; i++ {
				if board[b[i][0]] == "O" && board[b[i][1]] == "O" && board[b[i][2]] == " " {
					board[b[i][2]] = "O"
					oTurn = false
					break
				}
				if board[b[i][0]] == "O" && board[b[i][1]] == " " && board[b[i][2]] == "O" {
					board[b[i][1]] = "O"
					oTurn = false
					break
				}
				if board[b[i][0]] == " " && board[b[i][1]] == "O" && board[b[i][2]] == "O" {
					board[b[i][0]] = "O"
					oTurn = false
					break
				}
			}
		}

		// Does O need a simple block?
		if oTurn {
			for i := 0; i < 8; i++ {
				if board[b[i][0]] == "X" && board[b[i][1]] == "X" && board[b[i][2]] == " " {
					board[b[i][2]] = "O"
					oTurn = false
					break
				}
				if board[b[i][0]] == "X" && board[b[i][1]] == " " && board[b[i][2]] == "X" {
					board[b[i][1]] = "O"
					oTurn = false
					break
				}
				if board[b[i][0]] == " " && board[b[i][1]] == "X" && board[b[i][2]] == "X" {
					board[b[i][0]] = "O"
					oTurn = false
					break
				}
			}
		}

		// Does O need a corner block?
		if oTurn {
			if (board[1] == "X" && board[9] == "X" && board[5] == "O") ||
				(board[3] == "X" && board[7] == "X" && board[5] == "O") {
				for i := 0; i < 3; i++ {
					if board[side[i]] == " " {
						board[side[i]] = "O"
						oTurn = false
						break
					}
				}

			}

		}

		// Can O set up for a win?
		if oTurn {
			for i := 0; i < 8; i++ {
				if board[b[i][0]] == "O" && board[b[i][1]] == " " && board[b[i][2]] == " " {
					board[b[i][1]] = "O"
					oTurn = false
					break
				}
				if board[b[i][0]] == " " && board[b[i][1]] == "O" && board[b[i][2]] == " " {
					board[b[i][0]] = "O"
					oTurn = false
					break
				}
				if board[b[i][0]] == " " && board[b[i][1]] == " " && board[b[i][2]] == "0" {
					board[b[i][1]] = "O"
					oTurn = false
					break
				}
			}
		}

		// Is middle available?
		if oTurn {
			if board[5] == " " {
				board[5] = "O"
				oTurn = false
			}
		}

		// Is corner available?
		if oTurn {
			for i := 0; i <= 3; i++ {
				if board[corner[i]] == " " {
					board[corner[i]] = "O"
					oTurn = false
					break
				}
			}
		}

		// Make a random O move
		if oTurn {
			for i := 1; i <= 9; i++ {
				if board[i] == " " {
					board[i] = "O"
					oTurn = false
					break
				}
			}
		}

		// Did O win?
		if (board[1] == "O" && board[2] == "O" && board[3] == "O") ||
			(board[4] == "O" && board[5] == "O" && board[6] == "O") ||
			(board[7] == "O" && board[8] == "O" && board[9] == "O") ||
			(board[1] == "O" && board[4] == "O" && board[7] == "O") ||
			(board[2] == "O" && board[5] == "O" && board[8] == "O") ||
			(board[3] == "O" && board[6] == "O" && board[9] == "O") ||
			(board[1] == "O" && board[5] == "O" && board[9] == "O") ||
			(board[3] == "O" && board[5] == "O" && board[7] == "O") {
			showBoard(board)
			fmt.Println("***************")
			fmt.Println("**           **")
			fmt.Println("**  O WINS!  **")
			fmt.Println("**           **")
			fmt.Println("***************")
			fmt.Println(" ")
			break
		}

		// Refresh board
		showBoard(board)

	}

}
