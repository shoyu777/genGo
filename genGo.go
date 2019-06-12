package genGo

import (
	"fmt"
)

type Gengo struct {
	Kanji    string
	Hiragana string
	Romaji   string
}

var gengoList map[int]Gengo

func init() {
	g := Gengo{"a", "b", "c"}
	g1 := Gengo{"令和", "れいわ", "Reiwa"}
	fmt.Println(g)
	gengoList = make(map[int]Gengo)
	gengoList[1989] = g
	gengoList[2019] = g1
}

//Search to seach gengo from YYYY
func Search(i int) Gengo {
	fmt.Println("search")
	return gengoList[2019]
}
