package genGo

import (
	"fmt"
	"testing"
)

func ExampleNew() {
	g, err := New(2019)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g.Romaji)
	//output -> REIWA GAN NEN
}

func TestGengoSearch(t *testing.T) {
	gg, _ := New(2019)
	if gg.Hiragana != "れいわ がんねん" {
		t.Error("Expected れいわ がんねん got", gg.Hiragana)
	}
	if gg.Kanji != "令和 元年" {
		t.Error("Expected 令和 元年 got", gg.Kanji)
	}
	if gg.Romaji != "REIWA GAN NEN" {
		t.Error("Expected REIWA GAN NEN got", gg.Romaji)
	}

	g2, _ := New(1200)
	if g2.Hiragana != "しょうじ 2ねん" {
		t.Error("Expected しょうじ 2ねん got", g2.Hiragana)
	}
	if g2.Kanji != "正治 2年" {
		t.Error("Expected 正治 2年 got", g2.Kanji)
	}
	if g2.Romaji != "SHOJI 2 NEN" {
		t.Error("Expected SHOJI 2 NEN got", g2.Romaji)
	}

	g3, _ := New(645)
	if g3.Kanji != "大化 元年" {
		t.Error("Expected first Gengo got nil")
	}

	_, err := New(644)
	if err == nil {
		t.Error("Expeceted Out of Range.")
	}

}
