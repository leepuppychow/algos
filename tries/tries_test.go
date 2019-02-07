package tries

import (
	"testing"
)

func TestInsert(t *testing.T) {
	root := Node{
		Children: make(map[rune]*Node),
		IsWord:   false,
	}
	root.Insert("hello")
	root.Insert("hall")
	root.Insert("hear")

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
	if root.Children['h'].Children['e'].Children['a'].IsWord == false {
		t.Errorf("Insert failed")
	}
}
