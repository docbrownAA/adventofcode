package day1

import (
	"fmt"
	"strconv"
)

func MainDay1Part2() {
	tabStrings := readInputs()
	mapMesurement := getMap(tabStrings)
	nbIncreased := 0
	for i := 1; i < len(mapMesurement); i++ {
		if mapMesurement[string(65+i-1)] < mapMesurement[string(65+i)] {
			nbIncreased++
		}

	}
	fmt.Println("NbIncreased:", nbIncreased)
}

func getMap(tabString []string) map[string]int {
	mapMesurement := make(map[string]int)
	var currLetter int = 64
	for i := 0; i < len(tabString); i++ {
		value, err := strconv.Atoi(tabString[i])
		check(err)

		if i < 3 {
			switch i % 3 {
			case 0:
				currLetter++
				mapMesurement[string(currLetter)] = mapMesurement[string(currLetter)] + value

			case 1:
				mapMesurement[string(currLetter)] = mapMesurement[string(currLetter)] + value
				mapMesurement[string(currLetter+1)] = mapMesurement[string(currLetter+1)] + value
			case 2:
				mapMesurement[string(currLetter)] = mapMesurement[string(currLetter)] + value
				mapMesurement[string(currLetter+1)] = mapMesurement[string(currLetter+1)] + value
				mapMesurement[string(currLetter+2)] = mapMesurement[string(currLetter+2)] + value
			}

		} else {

			switch i % 3 {
			case 0:
				currLetter += 3
				mapMesurement[string(currLetter)] = mapMesurement[string(currLetter)] + value
				mapMesurement[string(currLetter-1)] = mapMesurement[string(currLetter-1)] + value
				mapMesurement[string(currLetter-2)] = mapMesurement[string(currLetter-2)] + value
			case 1:
				mapMesurement[string(currLetter)] = mapMesurement[string(currLetter)] + value
				mapMesurement[string(currLetter-1)] = mapMesurement[string(currLetter-1)] + value
				mapMesurement[string(currLetter+1)] = mapMesurement[string(currLetter+1)] + value
			case 2:
				mapMesurement[string(currLetter)] = mapMesurement[string(currLetter)] + value
				mapMesurement[string(currLetter+1)] = mapMesurement[string(currLetter+1)] + value
				mapMesurement[string(currLetter+2)] = mapMesurement[string(currLetter+2)] + value
			}

		}

	}
	return mapMesurement
}
