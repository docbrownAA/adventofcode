package tuto

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LesBoucles() {
	for i := 0; i < 100; i++ {
		fmt.Println(i+1, ") Je ne dois pas être un connard.")
	}
}

func LaBoucleDuVideur() {
	scanner := bufio.NewScanner(os.Stdin)
	var age int

	for age < 18 {
		fmt.Println("C'est quoi ton age ? ")
		scanner.Scan()

		age, _ = strconv.Atoi(scanner.Text())

	}
	fmt.Print("Bienvenue")
}
