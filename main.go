package main

import (
	"adventofcode/day3"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"tuto/structures"
)

func main() {
	//day3.MainDay3()
	day3.MainDay3Part2()
	//day1.MainDay1Part2()
	//lectureecriture.MainLectureEcriture()
	//goroutines.MainGoRoutines()
	//tpmoropionclasse.MainMorpion()
	/* fmt.Println("Je m'appelle", affichage.Nom)
	fmt.Println(affichage.AfficheSexe())
	fmt.Println(passion.AffichagePassion())

	fmt.Println("-----------------") */
	//gestionPersonnage()
	//erreurs.Main()
	//interfaces.Main()
	//sliceetmap.MainSlice()
	//sliceetmap.MainMap()
	//structures.Main()
	//pointeurs.MainPointeurs()
	//tpmorpion.Start()
	//videur()
	//testSwitch()
	//tuto.LaBoucleDuVideur()
	/* sum, length := fonctions.Addition(8, 2, 14, 17, 2, 8, 4, 12, 6, 10)
	fmt.Println("Somme: ", sum)
	fmt.Println("Longueur: ", length) */

	//fonctions.MainFonction()
	//tableaux.MainTableaux()

	/* var (
		vie     int     = 20
		nom     string  = "Default"
		vitesse float32 = 5.4
	)
	fmt.Println("Vie: ", vie)
	fmt.Println("Nom: ", nom)
	fmt.Println("vitesse: ", vitesse)

	/*
		Déclaration des variables dynamiques
	*/
	/*flt := 15.5
	in := 5
	st := "hello"
	bol := true

	fmt.Printf("Le type de la variable flt est %T\n", flt)
	fmt.Printf("Le type de la variable in est %T\n", in)
	fmt.Printf("Le type de la variable st est %T\n", st)
	fmt.Printf("Le type de la variable bol est %T\n", bol)

	var x int = 32
	var y float32 = 50.2

	fmt.Println("x + y = ", float32(x)+y)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ente un truc: ")
	scanner.Scan()
	entreeUtilisateur := scanner.Text()
	fmt.Println(entreeUtilisateur)

	fmt.Print("Un nombre entier cette fois: ")
	scanner.Scan()
	nbr, _ := strconv.Atoi(scanner.Text())
	fmt.Printf("res: %d\n", (nbr + 6)) */
}

func gestionPersonnage() {
	p1 := structures.New("barbare", 200, 20, false, [3]string{"épée", "bouclier", "armure"})
	p2 := structures.New("magicien", 100, 40, false, [3]string{"potions", "poisons", "bâton"})

	p1.Affichage()
	p2.Affichage()

	for !p1.GetMort() && !p2.GetMort() {
		if p1.GetVie() < 70 {
			p1.Soigner()
		}
		if p2.GetVie() < 50 {
			p2.Soigner()
		}
		p1.Attaquer(&p2)
		p2.SetMort()
		if !p2.GetMort() {
			p2.Attaquer(&p1)
			p1.SetMort()
		}

	}

	p1.Affichage()
	p2.Affichage()
}

func videur() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entre ton age: ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("N'importe quoi !")
		os.Exit(2)
	}

	fmt.Print("Ton prénom: ")
	scanner.Scan()
	prenom := scanner.Text()

	/*
		On a besoin de changer la graine (générateur de nombres pseudo-aléatoires) à chaque exécution sinon
		on obtient le même nombre aléatoire
	*/

	rand.Seed(time.Now().UnixNano())
	randomInt1 := rand.Intn(2)
	randomInt2 := rand.Intn(2)

	if age < 18 {
		fmt.Println("Dégage !")
	} else if prenom == "Gaël" || prenom == "gael" {
		fmt.Println("J'aime pas ce prénom, sors !")
	} else if age == 18 && randomInt1 == 1 {
		fmt.Println("T'as du bol, rentre")
	} else if randomInt2 == 0 {
		fmt.Println("Pas de bol, sors !")
	} else {
		fmt.Println("Tu peux rentrer")
	}
}

func testSwitch() {

	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("On est samedi !")
	case time.Sunday:
		fmt.Println("On est dimanche")
	default:
		fmt.Println("Au boulot !")
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Fais ton choix: ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Un entier stp !")
		os.Exit(2)
	}

	switch choix {
	case 0, 1:
		println("George bool !")
	case 7:
		println("William Van de Walle!")
	case 23:
		println("Pour certains, le nombre 23 est source de nombreux mystères, qu'en penses-tu Jim Carrey?")
	case 42:
		println("la réponse à la question ultime du sens de la vie!")
	case 666:
		println("Quand le diable veut une âme, le mal devient séduisant.")
	default:
		println("Mauvais choix !") // Valeur par défaut
	}
}
