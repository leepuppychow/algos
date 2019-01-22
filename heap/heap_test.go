package heap

import (
	"fmt"
	"testing"
)

// MIN HEAP PROPERTY:
// parent of any node must be <= current node

//					5
//				/   \
//			 8		 10
//			/ \    / \
//		12   9  11  14
//   /  \
//  25  16

// Indices: [0, 1,  2,  3, 4,  5,  6,  7,  8]
// Heap:    [5, 8, 10, 12, 9, 11, 14, 25, 16]

// Parent: Math.floor((i-1)/2)
// R child: 2*index + 2
// L child: 2*index + 1

func Setup() MinHeap {
	return MinHeap{5, 8, 10, 12, 9, 11, 14, 25, 16}
}

func TestMin(t *testing.T) {
	h := Setup()
	if h.Min() != 5 {
		t.Errorf("Min function failed")
	}
}

func TestLength(t *testing.T) {
	h := Setup()
	if h.Length() != 9 {
		t.Errorf("Length function failed")
	}
}

func TestInsert(t *testing.T) {
	h := Setup()
	h.Insert(3)

	if h.Min() != 3 || h.Length() != 10 {
		t.Errorf("Insert function failed")
	}

	h2 := Setup()
	h2.Insert(7)

	if h2.Min() != 5 || h2.Length() != 10 {
		t.Errorf("Insert function failed")
	}

	h3 := Setup()
	h3.Insert(9)

	if h3.Min() != 5 || h3.Length() != 10 {
		t.Errorf("Insert function failed")
	}
}

func TestPopMin(t *testing.T) {
	h := Setup()

	if h.PopMin() != 5 {
		t.Errorf("PopMin function failed")
	}
	if h.Length() != 8 && h.Min() != 8 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 7 && h.Min() != 9 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 6 && h.Min() != 10 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 5 && h.Min() != 11 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 4 && h.Min() != 12 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 3 && h.Min() != 14 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 2 && h.Min() != 16 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 1 && h.Min() != 25 {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)

	h.PopMin()
	if h.Length() != 0  {
		t.Errorf("PopMin function failed")
	}
	fmt.Println(h)
}
