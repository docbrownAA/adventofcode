package tpmorpion

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	var choix int
	var err, errplacement error
	var plateau = [3][3]string{{"[ ]", "[ ]", "[ ]"}, {"[ ]", "[ ]", "[ ]"}, {"[ ]", "[ ]", "[ ]"}}
	joueur := "O"
	gagnant, termine := checkWin(plateau)
	for !termine {
		fmt.Print("Joueur ", joueur, " Entre un nombre entre 1 et 9: ")
		scanner.Scan()
		choix, err = strconv.Atoi(scanner.Text())
		if err == nil {
			plateau, errplacement = play(plateau, choix, joueur)
			if errplacement != nil {
				fmt.Println(errplacement.Error())
			} else {
				if joueur == "O" {
					joueur = "X"
				} else {
					joueur = "O"
				}
				displayPlateau(plateau)
				gagnant, termine = checkWin(plateau)
			}
		} else {
			fmt.Println("Entre un nombre et pas autre chose")
		}
	}
	fmt.Println(gagnant)
	os.Exit(0)
}

func play(plateau [3][3]string, choix int, joueur string) ([3][3]string, error) {
	if choix < 1 || choix > 9 {
		return plateau, errors.New("entre un nombre compris entre 1 et 9")
	}
	x := (choix - 1) / 3
	y := (choix - 1) % 3
	if plateau[x][y] != "[X]" && plateau[x][y] != "[O]" {
		plateau[x][y] = "[" + joueur + "]"
	} else {
		return plateau, errors.New("cette case est déjà prise")
	}
	return plateau, nil
}

func fini(plateau [3][3]string) bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if plateau[x][y] != "[X]" && plateau[x][y] != "[O]" {
				return false
			}

		}
	}
	return true
}

func checkWin(plateau [3][3]string) (string, bool) {
	var firstValue, secValue, thirdValue string
	for x := 0; x < 3; x++ {
		firstValue = plateau[x][0]
		secValue = plateau[x][1]
		thirdValue = plateau[x][2]
		if firstValue == secValue && secValue == thirdValue && firstValue != "[ ]" {
			if firstValue == "[X]" {
				return "Joueur X à gagné", true
			} else {
				return "Joueur O à gagné", true
			}
		}

	}
	for y := 0; y < 3; y++ {
		firstValue = plateau[0][y]
		secValue = plateau[1][y]
		thirdValue = plateau[2][y]
		if firstValue == secValue && secValue == thirdValue && firstValue != "[ ]" {
			if firstValue == "[X]" {
				return "Joueur X à gagné", true
			} else {
				return "Joueur O à gagné", true
			}
		}

	}

	firstValue = plateau[0][0]
	secValue = plateau[1][1]
	thirdValue = plateau[2][2]
	if firstValue == secValue && secValue == thirdValue && firstValue != "[ ]" {
		if firstValue == "[X]" {
			return "Joueur X à gagné", true
		} else {
			return "Joueur O à gagné", true
		}
	}

	firstValue = plateau[2][0]
	secValue = plateau[1][1]
	thirdValue = plateau[0][2]
	if firstValue == secValue && secValue == thirdValue && firstValue != "[ ]" {
		if firstValue == "[X]" {
			return "Joueur X à gagné", true
		} else {
			return "Joueur O à gagné", true
		}
	}

	return "Egalité", fini(plateau)
}

func displayPlateau(plateau [3][3]string) {
	for x := 2; -1 < x; x-- {
		for y := 0; y < 3; y++ {
			fmt.Print(" ", plateau[x][y])
		}
		fmt.Println()
	}
	/* fmt.Print(" ", plateau[2][0], " ")
	fmt.Print(" ", plateau[2][1], " ")
	fmt.Println(" ", plateau[2][2], " ")

	fmt.Print(" ", plateau[1][0], " ")
	fmt.Print(" ", plateau[1][1], " ")
	fmt.Println(" ", plateau[1][2], " ")

	fmt.Print(" ", plateau[0][0], " ")
	fmt.Print(" ", plateau[0][1], " ")
	fmt.Println(" ", plateau[0][2], " ") */
}
