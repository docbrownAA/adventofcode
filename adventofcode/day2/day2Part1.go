package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var LINK_INPUT = "C:\\dev\\Workspace\\go\\src\\adventofcode\\day2\\input.txt"

func MainDay2Part1() {
	readInputs()
}

func readInputs() {
	file, err := os.Open(LINK_INPUT)
	check(err)
	defer file.Close()
	horizontal := 0
	depth := 0
	aim := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		command := str[0]
		value, err := strconv.Atoi(str[1])
		check(err)
		switch command {
		case "forward":
			horizontal += value
			depth = depth + aim*value
		case "down":
			//depth += value
			aim += value
		case "up":
			//depth -= value
			aim -= value
		}
		fmt.Println("----")
		fmt.Println("C:", str)
		fmt.Println("H:", horizontal)
		fmt.Println("D:", depth)
		fmt.Println("A:", aim)

	}

	result := horizontal * depth
	fmt.Println("----")
	fmt.Println("Total", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
