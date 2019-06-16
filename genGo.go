package genGo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
)

type gengo struct {
	StartYear int
	Kanji     string
	Romaji    string
	Hiragana  string
}

//GengoResult has the result of search with Kanji, Romaji, Hiragana characters.
type GengoResult struct {
	Kanji    string
	Romaji   string
	Hiragana string
}

var gengoList []gengo

//New search Gengo and return GengoResult
func New(year int) (GengoResult, error) {
	//year range checking
	err := yearRangeCheck(year)
	if err != nil {
		fmt.Println(err)
		g := GengoResult{"", "", ""}
		return g, err
	}

	var sr gengo
	sr = search(year)
	return setResult(year, sr), nil
}

func init() {
	gengoList = []gengo{}

	//get file path of 'genGo.go' and create a path of data source, and open.
	_, filename, _, _ := runtime.Caller(0)
	fr, err := os.Open(path.Join(path.Dir(filename), "era_names.csv"))
	failOnError(err)

	defer fr.Close()

	r := csv.NewReader(fr)
	r.Comma = ','

	rows, err := r.ReadAll()
	failOnError(err)

	//set Gengo struct to a list from csv file
	for _, v := range rows {
		i, _ := strconv.Atoi(v[0])
		g := gengo{i, v[1], v[2], v[3]}
		gengoList = append(gengoList, g)
	}

}

//search to seach gengo from YYYY using binary search.
func search(target int) gengo {
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
	return gengoList[0]
}

func setResult(year int, gg gengo) GengoResult {
	var gr GengoResult

	eraYear := year - gg.StartYear + 1
	if eraYear == 1 {
		gr.Hiragana = gg.Hiragana + " がんねん"
		gr.Kanji = gg.Kanji + " 元年"
		gr.Romaji = gg.Romaji + " GAN NEN"
	} else {
		gr.Hiragana = gg.Hiragana + " " + strconv.Itoa(eraYear) + "ねん"
		gr.Kanji = gg.Kanji + " " + strconv.Itoa(eraYear) + "年"
		gr.Romaji = gg.Romaji + " " + strconv.Itoa(eraYear) + " NEN"
	}
	return gr
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
