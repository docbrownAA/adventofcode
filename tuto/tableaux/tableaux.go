package tableaux

import "fmt"

func MainTableaux() {
	boucleSurTableau()
}

func boucleSurTableau() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	fmt.Println(jours[:])   //récupère tous les éléments
	fmt.Println(jours[3:])  // récupère tous les éléments à partir du 3e
	fmt.Println(jours[:3])  // récupère les 3 prémiers éléments
	fmt.Println(jours[1:3]) // récupère de l'index 1 à l'index 3 (exclus)
	for index, j := range jours {
		fmt.Println(j, "est le numéro", (index + 1), "de la semaine")
	}
}
