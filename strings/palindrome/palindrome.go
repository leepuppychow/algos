package palindrome

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Palindrome struct {
	Word  string
	Start int
	End   int
}

// Find longest palindromic substring
// This algorithm works... but could be refactored a lot. Probably wasting too much space for this problem.

func LongestPalindrome(s string) string {
	// Edge cases of strings of length 1 and 2
	if len(s) <= 1 {
		return s
	} else if len(s) == 2 && s[0] == s[1] {
		return s
	}

	palindromes := []Palindrome{}
	for i := 0; i <= len(s); i++ {
		if i == 0 { // skip index 0 because it will get checked on index 1
			continue
		} else if i == len(s)-1 { // stop, would be checked on index right before end of string
			break
		} else {
			AddPalindrome(s, i, &palindromes)
		}
	}
	fmt.Println("first set of palindromes:", palindromes)
	_, finalists, _ := PalindromeScan(s, palindromes, true)
	if len(finalists) > 0 {
		return FindLongest(finalists)
	}
	return string(s[0])
}

func AddPalindrome(s string, i int, palindromes *[]Palindrome) *[]Palindrome {
	if s[i+1] == s[i] {
		p := Palindrome{
			Word:  s[i : i+2],
			Start: i,
			End:   i + 1,
		}
		*palindromes = append(*palindromes, p)
	} else if s[i-1] == s[i] {
		p := Palindrome{
			Word:  s[i-1 : i+1],
			Start: i - 1,
			End:   i,
		}
		*palindromes = append(*palindromes, p)
	}

	if s[i+1] == s[i-1] {
		p := Palindrome{
			Word:  s[i-1 : i+2],
			Start: i - 1,
			End:   i + 1,
		}
		*palindromes = append(*palindromes, p)
	}
	return palindromes
}

func PalindromeScan(s string, palindromes []Palindrome, keepGoing bool) (string, []Palindrome, bool) {
	keepGoing = false
	newPalindromes := []Palindrome{}
	for _, pali := range palindromes {
		if pali.End+1 >= len(s) || pali.Start-1 < 0 {
			continue
		} else if s[pali.End+1] == s[pali.Start-1] { // Found a longer palindrome
			keepGoing = true
			p := Palindrome{
				Word:  s[pali.Start-1 : pali.End+2],
				Start: pali.Start - 1,
				End:   pali.End + 1,
			}
			newPalindromes = append(newPalindromes, p)
		}
	}
	if !keepGoing {
		return s, palindromes, false
	} else {
		fmt.Println("next set of palindromes", newPalindromes)
		return PalindromeScan(s, newPalindromes, keepGoing)
	}
}

func FindLongest(finalists []Palindrome) string {
	currentMax := len(finalists[0].Word)
	longest := finalists[0].Word
	for _, pali := range finalists {
		if len(pali.Word) > currentMax {
			longest = pali.Word
		}
	}
	return longest
}

// Check if string is a palindrome
func IsPalindrome(s string) bool {
	s = LetterAndNumOnly(strings.ToLower(s))
	isPalindrome := true
	for i, _ := range s {
		if s[i] != s[len(s)-1-i] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}

func LetterAndNumOnly(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Println("Error compiling regex", err)
	}
	return reg.ReplaceAllString(s, "")
}
