package sauvegarde

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"tuto/tpmoropionclasse/game"
)

type Save struct {
	filePath string
}

func New(filePath string) Save {

	save := Save{filePath: filePath}
	return save
}

func (s *Save) CheckExistingSave() bool {
	file, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	data := read(file.Name())
	defer file.Close()
	return len(data) > 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (s *Save) Write(text string) {
	file, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
	defer file.Close()
}

func read(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data)
}

func (s *Save) ExtractGame() game.Game {
	file, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	data := read(file.Name())
	var game game.Game
	json.Unmarshal([]byte(data), &game)
	defer file.Close()
	return game
}

func (s *Save) SaveGame(game *game.Game) {
	data, _ := json.Marshal(game)
	s.Write(string(data))
}
