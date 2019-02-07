package tries

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

func (n *Node) Insert(word string) {
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
	// fmt.Println(word, n)

	if len(word) == 0 {
		return
	} else if len(word) == 1 {
		n.IsWord = true
	}

	letter := rune(word[0])
	if _, ok := n.Children[letter]; ok {
		n.Children[letter].Insert(word[1:])
	} else {
		n.Children[letter] = &Node{
			Children: make(map[rune]*Node),
			IsWord:   false,
		}
		n.Children[letter].Insert(word[1:])
	}
}
