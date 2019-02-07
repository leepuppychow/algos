package tries

import (
	"testing"
)

func Setup() Node {
	root := Node{
		Children: make(map[rune]*Node),
		IsWord:   false,
	}
	root.Insert("hello")
	root.Insert("hall")
	root.Insert("hear")

	return root
}

func TestInsert(t *testing.T) {
	root := Setup()

	if root.Children['h'] == nil {
		t.Errorf("Insert failed")
	}
	if root.Children['h'].Children['e'] == nil {
		t.Errorf("Insert failed")
	}
	if root.Children['h'].Children['a'] == nil {
		t.Errorf("Insert failed")
	}
	if root.Children['h'].Children['l'] != nil {
		t.Errorf("Insert failed")
	}
	if root.Children['h'].Children['e'].IsWord == true {
		t.Errorf("Insert failed")
	}
	if root.Children['h'].Children['e'].Children['a'].Children['r'].IsWord == false {
		t.Errorf("Insert failed")
	}
}

func TestSearch(t *testing.T) {
	root := Setup()

	if !root.Search("hello") {
		t.Errorf("Search failed")
	}
	if !root.Search("hall") {
		t.Errorf("Search failed")
	}
	if !root.Search("hear") {
		t.Errorf("Search failed")
	}
	if root.Search("ball") {
		t.Errorf("Search failed")
	}
	if root.Search("hell") {
		t.Errorf("Search failed")
	}
}
