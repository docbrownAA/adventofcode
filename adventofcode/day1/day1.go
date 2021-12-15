package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var LINK_INPUT = "C:\\dev\\Workspace\\go\\src\\adventofcode\\day1\\input.txt"

func MainDay1() {
	tabStrings := readInputs()
	var nbIncreased int = 0
	var nbDecreased int = 0
	for i := 1; i < len(tabStrings); i++ {
		fmt.Println(tabStrings[i])
		prev, err := strconv.Atoi(tabStrings[i-1])
		check(err)
		curr, err := strconv.Atoi(tabStrings[i])
		check(err)
		fmt.Println(prev)
		if prev < curr {
			nbIncreased = nbIncreased + 1
		} else {
			nbDecreased++
		}
	}
	fmt.Println("NbIncreased :", nbIncreased)
	fmt.Println("NBDecreased :", nbDecreased)
}

func readInputs() []string {
	var tabString []string
	file, err := os.Open(LINK_INPUT)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		tabString = append(tabString, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tabString
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
