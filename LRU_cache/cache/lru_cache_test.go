package cache

import "testing"

func TestSet(t *testing.T) {
	lru := Constructor(3)
	lru.Set(1, 10)
	if len(lru.Cache) != 1 || lru.Cache[1].Value != 10 {
		t.Errorf("Set method failed, expect cache to have 1 item")
	}
	if lru.Head.Key != 1 || lru.Head.Value != 10 {
		t.Errorf("Set method failed, expect lru Head to have been set correctly")
	}
	if lru.Head != lru.Tail {
		t.Errorf("Set method failed, with 1 item the head and tail should be the same")
	}

	lru.Set(2, 20)
	if len(lru.Cache) != 2 || lru.Cache[2].Value != 20 {
		t.Errorf("Set method failed, expect cache to have 2 items")
	}
	if lru.Head.Key != 2 || lru.Head.Value != 20 {
		t.Errorf("Set method failed, expect lru Head to have been set correctly")
	}
	if lru.Head == lru.Tail || lru.Tail.Value != 10 {
		t.Errorf("Set method failed, with 2 items the head and tail should be different now")
	}

	actualListOrder := lru.Head.PrintList("")
	if actualListOrder != "21" {
		t.Errorf("Expect list order to be 21, got %s", actualListOrder)
	}

	lru.Set(3, 30)
	actualListOrder = lru.Head.PrintList("")
	if actualListOrder != "321" {
		t.Errorf("Expect list order to be 321, got %s", actualListOrder)
	}
}
