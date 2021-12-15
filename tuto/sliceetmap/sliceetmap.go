package sliceetmap

import "fmt"

func MainSlice() {
	var nombres = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(nombres)
	// ou
	var altNombre = make([]int, 5)
	fmt.Println(altNombre)

	// Ajout d'element
	var mois []string
	mois = append(mois, "Janvier")
	fmt.Println(mois)
	mois = append(mois, "Février")
	fmt.Println(mois)

	// Suppression d'un élément dans un slice
	mois = append(mois, "Mars", "Avril", "Mai", "Juin")
	fmt.Println(mois)

	indexASupprimer := 1

	mois = append(mois[:indexASupprimer], mois[(indexASupprimer+1):]...)
	fmt.Println(mois)

	// Copy d'une slice

	mois2 := make([]string, len(mois))

	copy(mois2, mois)

	fmt.Println(mois2)

}

func MainMap() {

	// 2 méthodes de création
	// 1)
	//var notes map[string]int

	// 2)
	//var notes2 = make(map[string]int)

	notes := map[string]int{"Sam": 20, "William": 18}
	// Ajout d'éléments
	notes["Gaël"] = 10
	notes["Christelle"] = 21

	// Lecture d'éléments
	fmt.Println("La note de Christelle est ", notes["Christelle"])
	fmt.Println("La note de Gaël est ", notes["Gaël"])
	fmt.Println("La note de Connard est ", notes["Connard"])

	//Lecture dans une boucle for
	for eleve := range notes {
		fmt.Println("La note de ", eleve, " est: ", notes[eleve])
	}

	// Suppression d'un élément
	delete(notes, "Gaël")

	for eleve := range notes {
		fmt.Println("La note de ", eleve, " est: ", notes[eleve])
	}
}
