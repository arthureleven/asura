package game

import (
	"encoding/json"
	"os"
)

type Rarity = int

type Class struct {
	Name          string `json:"name"`
	Rarity        Rarity `json:"rarity"`
	Disadvantages []int  `json:"disadvantages"`
}

var Classes []*Class

func Init() {
	if bytes, err := os.ReadFile("resources/classes.json"); err != nil {
		json.Unmarshal([]byte(bytes), &Classes)
	}
}
