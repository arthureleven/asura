package game

import (
	"asura/prototyping"
	"asura/schemas"
	"encoding/json"
	"math"
	"os"
)

type Rooster struct {
	Name          string   `json:"name"`
	Rarity        int      `json:"rarity"`
	Disadvantages []int    `json:"disadvantages"`
	Moves         []Skill  `json:"moves"`
	Sprites       []string `json:"sprites"`
}

type Skill struct {
	Name     string `json:"name"`
	Damage   [2]int `json:"damage"`
	Required int    `json:"required"`
}

var Roosters []*Rooster

func Init() {
	if bytes, err := os.ReadFile("resources/roosters.json"); err == nil {
		json.Unmarshal([]byte(bytes), &Roosters)
	}
}

func CalcLevel(xp int) int {
	return int(math.Floor(math.Sqrt(float64(xp)/30))) + 1
}

func CalcDamage(skill *Skill, rooster *schemas.Rooster) (min, max int) {
	ranges := prototyping.Map(skill.Damage[:], func(damage int) int {
		return damage + int(float64(damage)*0.15*float64(rooster.Resets))
	})

	return ranges[0], ranges[1]
}
