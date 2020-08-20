package main

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

func boardFull() bool {
	for _, v := range a {
		if v == " " {
			return false
		}
	}

	return true
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
