package plateau

import (
	"errors"
	"fmt"
	"strconv"
	"tuto/tpmoropionclasse/joueur"
)

type Plateau struct {
	Dimension int
	Cases     [][]string
}

func New(dimension int) Plateau {
	cases := make([][]string, dimension)
	for v := range cases {
		cases[v] = make([]string, dimension)
	}
	var nb int = 1
	for x := 0; x < dimension; x++ {
		for y := 0; y < dimension; y++ {
			cases[x][y] = "[" + strconv.Itoa(nb) + "]"
			nb++
		}
	}

	plateau := Plateau{dimension, cases}
	return plateau
}

func (p *Plateau) Affichage() {
	fmt.Println("------------")
	for x := p.Dimension - 1; -1 < x; x-- {
		for y := 0; y < p.Dimension; y++ {
			fmt.Print(p.Cases[x][y])
		}
		fmt.Println("")
	}
}

func (p *Plateau) AffichageCoordonee() {
	for i := 0; i < p.Dimension; i++ {
		for j := 0; j < p.Dimension; j++ {
			fmt.Print("[", i, ",", j, "]")
		}
		fmt.Println("")
	}
}

func (p *Plateau) AffichageReel() {
	for i := 0; i < p.Dimension; i++ {
		for j := 0; j < p.Dimension; j++ {
			fmt.Print(p.Cases[i][j])
		}
		fmt.Println("")
	}
}

func (p *Plateau) Remplir(choix int, j *joueur.Joueur) error {
	if choix < 1 || choix > p.Dimension*p.Dimension {
		return errors.New("entre un nombre compris entre 1 et 9")
	}
	x := (choix - 1) / p.Dimension
	y := (choix - 1) % p.Dimension
	if p.Cases[x][y] != "[X]" && p.Cases[x][y] != "[O]" {
		p.Cases[x][y] = "[" + j.Symbole + "]"
	} else {
		return errors.New("cette case est déjà prise")
	}
	return nil
}

func (p *Plateau) CheckWin() (bool, string) {
	symbole := ""
	win := true
	for i := 0; i < p.Dimension; i++ {
		for j := 1; j < p.Dimension; j++ {
			symbole = p.Cases[i][j]
			win = (p.Cases[i][j-1] == p.Cases[i][j] && p.Cases[i][j] != "[ ]")
			if !win {
				break
			}
		}
		if win {
			break
		}
	}
	if win {
		return win, symbole
	}
	for j := 0; j < p.Dimension; j++ {
		for i := 1; i < p.Dimension; i++ {
			symbole = p.Cases[i][j]
			win = (p.Cases[i-1][j] == p.Cases[i][j] && p.Cases[i][j] != "[ ]")
			if !win {
				break
			}
		}
		if win {
			break
		}
	}
	if win {
		return win, symbole
	}

	for i := 1; i < p.Dimension; i++ {
		symbole = p.Cases[i][i]
		win = (p.Cases[i-1][i-1] == p.Cases[i][i] && p.Cases[i][i] != "[ ]")
		if !win {
			break
		}
	}
	if win {
		return win, symbole
	}

	for i := 1; i < p.Dimension; i++ {
		symbole = p.Cases[i-1][p.Dimension-i]
		win = (p.Cases[i][p.Dimension-i-1] == symbole && symbole != "[ ]")
		if !win {
			break
		}
	}
	if win {
		return win, symbole
	}
	symbole = "[ ]"
	return win, symbole
}

func (p *Plateau) Fini() bool {
	for x := 0; x < p.Dimension; x++ {
		for y := 0; y < p.Dimension; y++ {
			if p.Cases[x][y] != "[X]" && p.Cases[x][y] != "[O]" {
				return false
			}

		}
	}
	return true
}
