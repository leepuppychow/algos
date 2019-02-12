package strings

/*

1. Iterate through the string
2. If the current letter matches the first letter of the pattern, then iterate through len of the pattern
3. If all characters of pattern match, push that index into the array
4. If not, continue with the next letter of the string
5. Break once you reach the end of the string (could stop at len(string) - len(pattern) to save some computation)

*/

// This would be ~O(n * m) with n = length of text and m = length of pattern

func IncludesNaive(text, pattern string) []int {
	matches := []int{}
	if pattern == "" || text == "" {
		return matches
	}

	for i, letter := range text {
		if letter == rune(pattern[0]) {
			match := true
			for j, patternLetter := range pattern {
				if patternLetter != rune(text[i+j]) {
					match = false
					break
				}
			}
			if match {
				matches = append(matches, i)
			}
		}
	}
	return matches
}
