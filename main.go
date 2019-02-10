package main

import (
	"fmt"
	"time"

	"github.com/leepuppychow/algos/arrays"
)

func timeTestTwoVars(fn func(a, b int) int, a, b int) {
	start := time.Now()
	fmt.Println(fn(a, b))
	fmt.Println(time.Since(start))
}

func timeTestOneVar(fn func(a int) int, a int) {
	start := time.Now()
	fmt.Println(fn(a))
	fmt.Println(time.Since(start))
}

func main() {
	fmt.Println("done", arrays.SingleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	// root := tries.Node{
	// 	Children: make(map[rune]*tries.Node),
	// 	IsWord:   false,
	// }
	// words := []string{
	// 	"hello",
	// 	"hall",
	// 	"hear",
	// 	"helmet",
	// 	"hearing",
	// 	"hand",
	// 	"hope",
	// 	"dog",
	// 	"dark",
	// 	"dragon",
	// 	"dracula",
	// 	"drag",
	// }
	// for _, w := range words {
	// 	root.Insert(w)
	// }
	// fmt.Println(root.Suggest(""))
}
