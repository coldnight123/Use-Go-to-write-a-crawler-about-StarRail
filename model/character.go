package model

import (
	"fmt"
)

type Character struct {
	Cid  int
	Name string
	//icon      string
	Attribute string
	Fate      string
	Star      string
	Camp      string
	Card      string
	Poster1   string
	Poster2   string
}

type CardSrc struct {
	SrcKey   string
	SrcValue string
}

func (char Character) CharacterPrint() {
	format :=
		"CID: %d\n" +
			"Name: %s\n" +
			// "Icon: %s\n" +
			"Attribute: %s\n" +
			"Fate: %s\n" +
			"Star: %s\n" +
			"Camp: %s\n" +
			"Card: %s\n" +
			"Poster1: %s\n" +
			"Poster2: %s\n"
	fmt.Printf(format, char.Cid, char.Name, char.Attribute, char.Fate, char.Star, char.Camp, char.Card, char.Poster1, char.Poster2)
}
