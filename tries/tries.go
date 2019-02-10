package tries

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

func (current *Node) Suggest(substring string) []string {
	for _, letter := range substring {
		current = current.Children[letter]
		if current == nil { // in case there are no matches to substring
			return []string{}
		}
	}
	words := []string{}
	current.FindWords(substring, &words)
	return words
}

func (n *Node) FindWords(currentWord string, words *[]string) {
	if len(n.Children) == 0 || n.IsWord {
		*words = append(*words, currentWord)
	}
	for letter, child := range n.Children {
		child.FindWords(currentWord+string(letter), words)
	}
}

// bufio.NewReader(os.Stdin).ReadBytes('\n')

// If deleting a word that is not defined by a leaf node, then just toggle the IsWord to false,
// If deleting a word that IS a leaf, then need to delete that key from the Children map 
//	In Go, you can do:  delete(mapName, "key")
func (n *Node) Delete(word string) { // make sure to prune tree as well

}

// A select function (weighting each substring and its top selected words)
// Array from suggest should then be ordered by the top selected choices 

func (n *Node) Select(substring, word string) {

}
