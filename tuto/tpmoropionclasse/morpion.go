package tpmoropionclasse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tuto/tpmoropionclasse/game"
	"tuto/tpmoropionclasse/joueur"
	"tuto/tpmoropionclasse/plateau"
	"tuto/tpmoropionclasse/sauvegarde"
)

var savePath string = "C:\\dev\\Workspace\\go\\src\\tuto\\tpmoropionclasse\\sauvegarde\\jeu.json"

func MainMorpion() {
	sauvegarde := sauvegarde.New(savePath)
	scanner := bufio.NewScanner(os.Stdin)

	win := false
	symbole := ""
	var partie game.Game
	var plateauJeu plateau.Plateau
	var joueur1 joueur.Joueur
	var joueur2 joueur.Joueur
	if sauvegarde.CheckExistingSave() {
		partie = sauvegarde.ExtractGame()
		plateauJeu = *partie.Plateau
		joueur1 = *partie.Joueur1
		joueur2 = *partie.Joueur2
	} else {
		fmt.Println("Nouvelle Partie")
		//fmt.Print("Joueur 1, entre ton nom:")
		//scanner.Scan()
		//nom := scanner.Text()
		nom := "Joueur1"
		joueur1 = joueur.New(nom, "O")
		joueur1.Affichage()
		//fmt.Print("Joueur 2, entre ton nom:")
		//scanner.Scan()
		nom = "Joueur2"
		joueur2 = joueur.New(nom, "X")
		joueur2.AJoue = true
		joueur2.Affichage()
		plateauJeu = plateau.New(3)
		partie = game.Init(plateauJeu, &joueur1, &joueur2)
	}

	for !plateauJeu.Fini() && !win {
		plateauJeu.Affichage()

		var errPlacement error
		if !joueur1.AJoue {
			fmt.Print("A ", joueur1.Nom, " de jouer:")

			scanner.Scan()
			choix, err := strconv.Atoi(scanner.Text())
			if err == nil {
				errPlacement = plateauJeu.Remplir(choix, &joueur1)
				if errPlacement != nil {
					fmt.Println(errPlacement)
					continue
				}
			} else {
				if strings.ToLower(scanner.Text()) == "exit" || strings.ToLower(scanner.Text()) == "x" {
					os.Exit(0)
				}
				fmt.Println("Entre un nombre et pas autre chose !")
				continue
			}
		}
		if !joueur2.AJoue {
			fmt.Print("A ", joueur2.Nom, " de jouer:")

			scanner.Scan()
			choix, err := strconv.Atoi(scanner.Text())
			if err == nil {
				errPlacement = plateauJeu.Remplir(choix, &joueur2)
				if errPlacement != nil {
					fmt.Println(errPlacement)
					continue
				}
			} else {
				if strings.ToLower(scanner.Text()) == "exit" || strings.ToLower(scanner.Text()) == "x" {
					os.Exit(0)
				}
				fmt.Println("Entre un nombre et pas autre chose !")
				continue
			}
		}
		joueur2.AJoue = !joueur2.AJoue
		joueur1.AJoue = !joueur1.AJoue
		win, symbole = plateauJeu.CheckWin()
		sauvegarde.SaveGame(&partie)

	}
	plateauJeu.Affichage()
	symbole = symbole[1 : len(symbole)-1]
	fmt.Println("Win:", win)
	if symbole == joueur1.Symbole {
		fmt.Println("Gagnant:", joueur1.Nom)
	} else if symbole == joueur2.Symbole {
		fmt.Println("Gagnant:", joueur2.Nom)
	} else {
		fmt.Println("pas de gagant")
	}
}
