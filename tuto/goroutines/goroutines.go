package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Instanciation de la structure WaitGroup

func MainGoRoutines() {
	debut := time.Now()

	wg.Add(1)
	go run("Gaël")
	wg.Add(1)
	go run("Christelle")
	wg.Add(1)
	go run("Sam")

	wg.Wait()
	fin := time.Now()
	fmt.Println(fin.Sub(debut))
}

func run(name string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, ": ", i)
	}
}
