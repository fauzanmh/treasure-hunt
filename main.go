package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	// check os
	currentOs := runtime.GOOS
	if currentOs != "linux" && currentOs != "windows" {
		println("Your OS is unsupported for running this program!")
		os.Exit(0)
	}

	// initialize variable clear screen
	clear = make(map[string]func())
	// set linux to variable clear screen
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	// set windows to variable clear screen
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	// print title a program
	fmt.Print("\n")
	fmt.Println("=======================================")
	fmt.Println("Treasunt Hunt")
	fmt.Println("Find Treasure at any of the Coordinates")
	fmt.Println("=======================================")
	fmt.Print("\n")

	// layout probable of treasure
	probableOfTreasure()

	// print a tool
	fmt.Print("\n")
	fmt.Println("======================================================")
	fmt.Println("Press S to start treasure hunt")
	fmt.Println("Press any key other than the S key to exit the program")
	fmt.Println("======================================================")

	// read typed
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	// check typed
	switch text {
	case "s", "S":
		// start treasure hunt
		clearScreen()
		treasureHunt()
	}

	os.Exit(0)
}

func probableOfTreasure() {
	// layout probable of treasure
	layoutProbableTreasure := [6][8]bool{
		{false, false, false, false, false, false, false, false},
		{false, true, true, true, true, true, true, false},
		{false, true, false, false, false, true, true, false},
		{false, true, true, true, false, true, false, false},
		{false, true, false, true, true, true, true, false},
		{false, false, false, false, false, false, false, false},
	}

	// first probable coordinates of treasure
	firstProbableCoordinateX := 4
	firstProbableCoordinateY := 6

	// second probable coordinates of treasure
	secondProbableCoordinateX := 1
	secondProbableCoordinateY := 5

	// second probable coordinates of treasure
	thirdProbableCoordinateX := 1
	thirdProbableCoordinateY := 1

	// print layout treasure hunt
	for idx := range layoutProbableTreasure {
		for idx2 := range layoutProbableTreasure[idx] {
			isFirstProbableCoordinate := idx == firstProbableCoordinateX && idx2 == firstProbableCoordinateY
			isSecondProbableCoordinate := idx == secondProbableCoordinateX && idx2 == secondProbableCoordinateY
			isThirdProbableCoordinate := idx == thirdProbableCoordinateX && idx2 == thirdProbableCoordinateY

			if !layoutProbableTreasure[idx][idx2] {
				fmt.Print(" #")
				continue
			} else if isFirstProbableCoordinate || isSecondProbableCoordinate || isThirdProbableCoordinate {
				fmt.Print(" $")
				continue
			}

			fmt.Print(" .")
		}

		fmt.Println("")
	}
}

func treasureHunt() {
	// layout treasure hunt
	layoutTreasureHunt := [6][8]bool{
		{false, false, false, false, false, false, false, false},
		{false, true, true, true, true, true, true, false},
		{false, true, false, false, false, true, true, false},
		{false, true, true, true, false, true, false, false},
		{false, true, false, true, true, true, true, false},
		{false, false, false, false, false, false, false, false},
	}

	isFounded := false
	isObstacle := false

	// start position in treasure hunt
	positionX := 4
	positionY := 1

	// start temp position for current position in treasure hunt
	tempPositionX := 4
	tempPositionY := 1

	// treasure position in treasure hunt
	rand.Seed(time.Now().UnixNano())
	randCoordinate := rand.Intn(3)
	treasureCoordinates := [][]int{{4, 6}, {1, 5}, {1, 1}}
	treasureCoordinate := treasureCoordinates[randCoordinate]

	for {
		// print layout treasure hunt
		for idx := range layoutTreasureHunt {
			for idx2 := range layoutTreasureHunt[idx] {
				// check current position
				if positionX == treasureCoordinate[0] && positionY == treasureCoordinate[1] && !isFounded {
					isFounded = true
					fmt.Print(" #")
					continue
				} else if idx == positionX && idx2 == positionY {
					if !layoutTreasureHunt[idx][idx2] {
						isObstacle = true
						break
					}

					if isFounded {
						fmt.Print(" $")
					} else {
						fmt.Print(" X")
					}

					tempPositionX = positionX
					tempPositionY = positionY
					continue
				} else if !layoutTreasureHunt[idx][idx2] {
					fmt.Print(" #")
					continue
				}

				fmt.Print(" .")
			}

			fmt.Println("")
		}

		// handle if moving to obstacle
		if isObstacle {
			positionX = tempPositionX
			positionY = tempPositionY
			isObstacle = false
			clearScreen()
			continue
		}

		// print text if treasure found
		if isFounded {
			fmt.Print("\n")
			fmt.Println("=================")
			fmt.Println("Treasure Found")
			fmt.Println("=================")
			fmt.Print("\n")
			os.Exit(0)
		}

		// print a tool
		fmt.Print("\n")
		fmt.Println("======================================")
		fmt.Println("Press W for moving up/north step(s)")
		fmt.Println("Press A for moving left/east step(s)")
		fmt.Println("Press S for moving down/south step(s)")
		fmt.Println("Press D for moving right/west step(s)")
		fmt.Println("Press C to exit the program")
		fmt.Println("======================================")

		// read typed
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\n", "", -1)

		// check typed
		switch text {
		case "w", "W":
			// moving up/north step(s)
			positionX -= 1
		case "a", "A":
			// moving left/east step(s)
			positionY -= 1
		case "s", "S":
			// moving down/south step(s)
			positionX += 1
		case "d", "D":
			// moving rigth/west step(s)
			positionY += 1
		case "c", "C":
			os.Exit(0)
		}

		clearScreen()
	}
}

// func to clear screen
func clearScreen() {
	exec, support := clear[runtime.GOOS]
	if support {
		exec()
	} else {
		fmt.Println("Your platform is unsupported")
		os.Exit(0)
	}
}
