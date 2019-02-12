package tries

import (
	"testing"
)

func Setup() Node {
	root := Node{
		Children: make(map[rune]*Node),
		IsWord:   false,
	}
	words := []string{
		"hello",
		"hall",
		"hear",
		"helmet",
		"hearing",
		"hand",
		"hope",
		"dog",
		"dark",
		"dragon",
		"dracula",
		"drag",
	}
	for _, w := range words {
		root.Insert(w)
	}

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

func TestSuggestRegular(t *testing.T) {
	root := Setup()

	words := root.Suggest("dra")
	expected := map[string]bool{
		"dragon":  true,
		"dracula": true,
		"drag":    true,
	}
	for _, word := range words {
		if !expected[word] {
			t.Errorf("Suggest failed")
		}
	}
	if len(words) != 3 {
		t.Errorf("Suggest failed")
	}

	words = root.Suggest("hear")
	expected = map[string]bool{
		"hear":    true,
		"hearing": true,
	}
	for _, word := range words {
		if !expected[word] {
			t.Errorf("Suggest failed")
		}
	}
	if len(words) != 2 {
		t.Errorf("Suggest failed")
	}

	words = root.Suggest("h")
	expected = map[string]bool{
		"hello":   true,
		"hall":    true,
		"hear":    true,
		"helmet":  true,
		"hearing": true,
		"hand":    true,
		"hope":    true,
	}
	for _, word := range words {
		if !expected[word] {
			t.Errorf("Suggest failed")
		}
	}
	if len(words) != 7 {
		t.Errorf("Suggest failed")
	}
}

func TestSuggestEdgeCases(t *testing.T) {
	root := Setup()

	words := root.Suggest("hearing")
	expected := map[string]bool{
		"hearing": true,
	}
	for _, word := range words {
		if !expected[word] {
			t.Errorf("Suggest failed")
		}
	}
	if len(words) != 1 {
		t.Errorf("Suggest failed")
	}

	words = root.Suggest("hearings")
	if len(words) != 0 {
		t.Errorf("Suggest failed")
	}

	words = root.Suggest("yo")
	if len(words) != 0 {
		t.Errorf("Suggest failed")
	}
	words = root.Suggest("")
	if len(words) != 12 {
		t.Errorf("Suggest failed")
	}
}

func TestDeleteWordNonLeaf(t *testing.T) {
	root := Setup()
	if !root.Search("hear") {
		t.Errorf("Delete Word failed")
	}
	root.Delete("hear", nil)
	if root.Search("hear") {
		t.Errorf("Delete Word failed")
	}
}

// Should also prune any children nodes that are no longer needed
func TestDeleteWordLeaf(t *testing.T) {
	root := Setup()
	if !root.Search("hearing") {
		t.Errorf("Delete Word failed")
	}
	root.Delete("hearing", nil)
	if root.Search("hearing") {
		t.Errorf("Delete Word failed")
	}
}
