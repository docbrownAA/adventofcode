package fonctions

import (
	"fmt"
	"math"
)

func Addition(param ...int) (int, int) {
	total := 0

	for _, value := range param {
		total += value
	}

	return total, len(param)
}

func MainFonction() {
	racineCarree := func(x float64) float64 { return math.Sqrt(x) }
	RajouterDix(9, racineCarree)

	RajouterDix(5, func(x float64) float64 { return math.Pow(x, 2) })
}

func RajouterDix(a float64, fAnonyme func(float64) float64) {
	operation := fAnonyme(a)
	result := operation + 10
	fmt.Println(result)
}
