package day3

import (
	"fmt"
	"math"
	"strconv"
)

func MainDay3Part2() {
	tabStrings := readInputs()
	resultOxy := getOxygen(tabStrings, 0)
	resultCO2 := getCO2(tabStrings, 0)
	fmt.Println(resultOxy)
	fmt.Println(resultCO2)

	fmt.Println("Result:", resultCO2*resultOxy)

}

func getCO2(tabStrings [][]string, j int) int {
	CO2 := 0
	var tmpTabStrings [][]string
	var count int

	for i := 0; i < len(tabStrings); i++ {
		value, err := strconv.Atoi(tabStrings[i][j])
		check(err)
		count += value
	}
	count = len(tabStrings) - count
	fmt.Println(count)
	if float64(count) > float64(len(tabStrings))/float64(2) {
		for i := 0; i < len(tabStrings); i++ {
			if tabStrings[i][j] == "1" {
				tmpTabStrings = append(tmpTabStrings, tabStrings[i])
			}
		}
	} else {
		for i := 0; i < len(tabStrings); i++ {
			if tabStrings[i][j] == "0" {
				tmpTabStrings = append(tmpTabStrings, tabStrings[i])
			}
		}
	}
	//fmt.Println(tmpTabStrings)
	if len(tmpTabStrings) == 1 {
		for i := len(tmpTabStrings[0]) - 1; i >= 0; i-- {
			value, err := strconv.Atoi(tmpTabStrings[0][i])
			check(err)
			CO2 = CO2 + (value * int(math.Pow(2, float64(len(tmpTabStrings[0])-1-i))))
		}
		return CO2
	} else {
		newJ := j + 1
		return getCO2(tmpTabStrings, newJ)
	}
}

func getOxygen(tabStrings [][]string, j int) int {
	oxygen := 0
	var tmpTabStrings [][]string
	var count int

	for i := 0; i < len(tabStrings); i++ {
		value, err := strconv.Atoi(tabStrings[i][j])
		check(err)
		count += value
	}
	if float64(count) >= float64(len(tabStrings))/float64(2) {
		for i := 0; i < len(tabStrings); i++ {
			if tabStrings[i][j] == "1" {
				tmpTabStrings = append(tmpTabStrings, tabStrings[i])
			}
		}
	} else {
		for i := 0; i < len(tabStrings); i++ {
			if tabStrings[i][j] == "0" {
				tmpTabStrings = append(tmpTabStrings, tabStrings[i])
			}
		}
	}
	//fmt.Println(tmpTabStrings)
	if len(tmpTabStrings) == 1 {
		for i := len(tmpTabStrings[0]) - 1; i >= 0; i-- {
			value, err := strconv.Atoi(tmpTabStrings[0][i])
			check(err)
			oxygen = oxygen + (value * int(math.Pow(2, float64(len(tmpTabStrings[0])-1-i))))
		}
		return oxygen
	} else {
		newJ := j + 1
		return getOxygen(tmpTabStrings, newJ)
	}

}
