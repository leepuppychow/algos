package tries

// bufio.NewReader(os.Stdin).ReadBytes('\n')
// fmt.Println(word, n)

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

func (n *Node) Insert(word string) {
	if len(word) == 0 {
		n.IsWord = true
		return
	}
	letter := rune(word[0])
	if _, ok := n.Children[letter]; !ok {
		n.Children[letter] = &Node{
			Children: make(map[rune]*Node),
			IsWord:   false,
		}
	}
	n.Children[letter].Insert(word[1:])
}

func (n *Node) Search(word string) bool {
	if len(word) == 0 && n.IsWord {
		return true
	} else if len(word) == 0 && !n.IsWord {
		return false
	}

	letter := rune(word[0])
	if _, ok := n.Children[letter]; ok {
		return n.Children[letter].Search(word[1:])
	} else {
		return false
	}
}

// func (n *Node) Suggest(substring string) []string {

// }
