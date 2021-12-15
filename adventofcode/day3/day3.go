package day3

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var LINK_INPUT = "C:\\dev\\Workspace\\go\\src\\adventofcode\\day3\\input.txt"

func MainDay3() {
	tabStrings := readInputs()
	result := getGammaEpsilon(tabStrings)
	fmt.Println("result:", result)
}

func readInputs() [][]string {
	var tabString [][]string
	file, err := os.Open(LINK_INPUT)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		tabString = append(tabString, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tabString
}

func getGammaEpsilon(tabStrings [][]string) int {
	gamma := 0
	espsilon := 0
	var count []int
	for i := 0; i < len(tabStrings[0]); i++ {
		count = append(count, 0)
	}
	for i := 0; i < len(tabStrings); i++ {
		for j := 0; j < len(count); j++ {
			value, err := strconv.Atoi(tabStrings[i][j])
			check(err)
			count[j] += value
		}
	}
	fmt.Println("count:", count)
	var strGamma []int
	var strEpsilon []int
	for i := 0; i < len(count); i++ {
		if count[i] > len(tabStrings)/2 {
			strGamma = append(strGamma, 1)
			strEpsilon = append(strEpsilon, 0)
		} else {
			strEpsilon = append(strEpsilon, 1)
			strGamma = append(strGamma, 0)
		}
	}

	for i := len(strGamma) - 1; i >= 0; i-- {
		gamma = gamma + (strGamma[i] * int(math.Pow(2, float64(len(strGamma)-1-i))))
	}
	for i := len(strEpsilon) - 1; i >= 0; i-- {
		espsilon = espsilon + (strEpsilon[i] * int(math.Pow(2, float64(len(strEpsilon)-1-i))))
	}
	fmt.Println(strGamma)
	fmt.Println("gamma", gamma)
	fmt.Println(strEpsilon)
	fmt.Println("epsilon", espsilon)
	return gamma * espsilon

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
