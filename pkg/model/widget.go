package model

import (
	"fmt"
	"math/rand"
	"time"
)

func getGems() []string {
	return []string{
		"Diamond",
		"Crystal",
		"Morion",
		"Azore",
		"Sapphire",
		"Cobalt",
		"Aquamarine",
		"Montana",
		"Turquoise",
		"Lime",
		"Erinite",
		"Emerald",
		"Turmaline",
		"Jonquil",
		"Olivine",
		"Topaz",
		"Citrine",
		"Sun",
		"Quartz",
		"Opal",
		"Alabaster",
		"Rose",
		"Burgundy",
		"Siam",
		"Ruby",
		"Amethyst",
		"Violet",
		"Lilac",
	}
}

func getColor() []string {
	return []string{
		"Blue",
		"Aqua",
		"Red",
		"Green",
		"Orange",
		"Yellow",
		"Black",
		"Violet",
		"Brown",
		"Crimson",
		"Gray",
		"Cyan",
		"Magenta",
		"White",
		"Gold",
		"Pink",
		"Lavender",
	}
}

func getThings() []string {
	return []string{
		"beard",
		"finger",
		"hand",
		"toe",
		"stalk",
		"hair",
		"vine",
		"street",
		"son",
		"brook",
		"river",
		"lake",
		"stone",
		"ship",
	}
}

func getRandomFromList(list []string) string {
	// rand.Seed(time.Now().UnixNano())
	minVal := 0
	val := rand.Intn(len(list)-1) + minVal
	return list[val]
}

func NewRandomWidget() Widget {
	name := fmt.Sprintf("%v-%v-%v", getRandomFromList(getColor()), getRandomFromList(getGems()), getRandomFromList(getThings()))
	return Widget{
		Name:        name,
		Description: fmt.Sprintf("The %v widget", name),
		Count:       rand.Intn(1000-1) + 1,
		Creator:     "Admin",
		CreatedAt:   time.Now(),
	}
}

func GetSeedData(total int) []Widget {
	var widgets []Widget
	for i := 0; i < total; i++ {
		widgets = append(widgets, NewRandomWidget())
	}
	return widgets
}
