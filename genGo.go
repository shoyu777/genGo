package genGo

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Gengo struct {
	StartYear int
	Kanji     string
	Romaji    string
	Hiragana  string
}

var gengoList []Gengo

func init() {
	gengoList = []Gengo{}

	fr, err := os.Open("./src/mypkg/genGo/era_names.csv")
	failOnError(err)

	defer fr.Close()

	r := csv.NewReader(fr)
	r.Comma = ','

	rows, err := r.ReadAll()
	failOnError(err)

	//set Gengo struct to a list from csv file
	for _, v := range rows {
		i, _ := strconv.Atoi(v[0])
		g := Gengo{i, v[1], v[2], v[3]}
		gengoList = append(gengoList, g)
	}

}

//Search to seach gengo from YYYY
func Search(target int) Gengo {
	head := 0
	tail := len(gengoList) - 1
	center := 0

	if gengoList[tail].StartYear <= target {
		return gengoList[tail]
	} else {
		for head <= tail {
			center = (head + tail) / 2

			if gengoList[center].StartYear <= target && target < gengoList[center+1].StartYear {
				return gengoList[center]
			} else if gengoList[center].StartYear < target {
				head = center + 1
			} else {
				tail = center - 1
			}
		}
	}

	return gengoList[1]
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
