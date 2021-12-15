package erreurs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func verificationDivisionParZero(nbr float64) (float64, error) {
	if nbr == 0 {
		return 0, errors.New("erreur: Il est impossible de diviser par 0")
	} else {
		return nbr, nil
	}
}

func division() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Entrez un nombre: ")
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Il faut saisir un nombre et pas autre chose.")
		} else {
			nbrDiv, errDiv := verificationDivisionParZero(float64(nbr))
			if errDiv != nil {
				panic(errDiv)
			} else {
				fmt.Println("Le résultat: ", 1000/nbrDiv)
				break
			}
		}

	}

}

func Main() {
	division()
	fmt.Println("Fin")
}
