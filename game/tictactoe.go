package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a [9]string
var turn = 1

// Point an x,y coordinate pair
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// PlaceValue adds a X or O to the tictactoe board
func PlaceValue(x int, y int) bool {
	if CheckWin() {
		return false
	}

	var ind = y*3 + x

	if ind >= 0 && ind < len(a) && a[ind] == " " {
		if turn%2 == 0 {
			a[ind] = "O"
		} else {
			a[ind] = "X"
		}

		turn++

		return true
	}

	return false
}

// BoardFull determines if there are any available tiles for a letter to be placed
func BoardFull() bool {
	for _, v := range a {
		if v == " " {
			return false
		}
	}

	return true
}

func getInd(x int, y int) int {
	var ind = y*3 + x

	if ind > len(a) {
		return -1
	}

	return ind
}

// GetWinner returns the character of the winning player.
func GetWinner() string {
	if turn%2 == 0 {
		return "X"
	}

	return "O"
}

// CheckWin determines if there is a winner
func CheckWin() bool {
	// Horizontal victory checks
	for i := 0; i < 3; i++ {
		col1 := a[getInd(0, i)]

		if col1 == " " {
			continue
		}

		col2 := a[getInd(1, i)]
		col3 := a[getInd(2, i)]

		if col1 == col2 && col2 == col3 {
			return true
		}
	}

	// Vertical victory checks
	for i := 0; i < 3; i++ {
		row1 := a[getInd(i, 0)]

		if row1 == " " {
			continue
		}

		row2 := a[getInd(i, 1)]
		row3 := a[getInd(i, 2)]

		if row1 == row2 && row2 == row3 {
			return true
		}
	}

	// Diagonals
	if a[0] != " " && a[0] == a[4] && a[0] == a[8] {
		return true
	}

	if a[2] != " " && a[2] == a[4] && a[2] == a[6] {
		return true
	}

	return false
}

// Reset the tictactoe game to a blank board
func Reset() {
	for i := range a {
		a[i] = " "
	}

	turn = 1
	fmt.Println()
	fmt.Println()
}

func clearDraw() {
	fmt.Printf("\nTurn %d:\n", turn)
}

// Draw draws the current state of the tictactoe board to stdout
func Draw() {
	clearDraw()

	for j := 0; j < 3; j++ {
		if j > 0 {
			fmt.Println("_____")
		}
		fmt.Printf("%v|%v|%v\n", a[j*3], a[j*3+1], a[j*3+2])
	}
}

func askInput() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	s := strings.Split(text, ",")

	var x, _ = strconv.Atoi(strings.Trim(s[0], " "))
	var y, _ = strconv.Atoi(strings.Trim(s[1], " \n"))

	return x, y
}
