package main

import "fmt"

var a [9]string
var turn = 1

func placeValue(x int, y int) bool {
	var ind = y*3 + x

	if ind >= 0 && ind < len(a) && a[ind] == " " {
		if turn%2 == 0 {
			a[ind] = "O"
		} else {
			a[ind] = "X"
		}

		turn++

		return true
	} else {
		return false
	}
}

func boardFull() bool {
	return false
}

func reset() {
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

func draw() {
	clearDraw()

	for j := 0; j < 3; j++ {
		if j > 0 {
			fmt.Println("_____")
		}
		fmt.Printf("%v|%v|%v\n", a[j*3], a[j*3+1], a[j*3+2])
	}
}

func main() {
	reset()

	draw()
	placeValue(1, 1)
	placeValue(1, 0)
	placeValue(2, 0)
	placeValue(0, 2)
	draw()
	placeValue(2, 1)
	placeValue(0, 1)
	placeValue(2, 2)

	draw()

	reset()

	draw()
}
