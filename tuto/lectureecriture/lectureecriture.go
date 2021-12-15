package lectureecriture

import (
	"fmt"
	"io/ioutil"
	"os"
)

var filePath string = "C:\\dev\\Workspace\\go\\src\\tuto\\lectureecriture\\test.txt"

func MainLectureEcriture() {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	check(err)

	write("Test\n", file)

	data := read(file.Name())

	fmt.Println(string(data))
	defer file.Close()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(text string, file *os.File) {
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
}

func read(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data)
}
