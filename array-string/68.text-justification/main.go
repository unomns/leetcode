package main

import (
	"fmt"
	"strings"
)

/**
Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible.
If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:
	- A word is defined as a character sequence consisting of non-space characters only.
	- Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
	- The input array words contains at least one word.

Example 1:
	Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
	Output:
	[
	"This    is    an",
	"example  of text",
	"justification.  "
	]

Example 2:
	Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
	Output:
	[
	"What   must   be",
	"acknowledgment  ",
	"shall be        "
	]
	Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
	Note that the second line is also left-justified because it contains only one word.

Example 3:
	Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
	Output:
	[
	"Science  is  what we",
	"understand      well",
	"enough to explain to",
	"a  computer.  Art is",
	"everything  else  we",
	"do                  "
	]


Constraints:
	1 <= words.length <= 300
	1 <= words[i].length <= 20
	words[i] consists of only English letters and symbols.
	1 <= maxWidth <= 100
	words[i].length <= maxWidth
*/

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	// fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	// fmt.Println()
	// fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	// fmt.Println()
	// fullJustify([]string{"Listen", "to", "many,", "speak", "to", "a", "few."}, 6)
	// fullJustify([]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}, 16)
}

func fullJustify(words []string, maxWidth int) []string {
	idxs := []int{0} // how many words in each row

	length := 0

	for i, word := range words {
		wordLength := len(word)
		if length+wordLength < maxWidth {
			length += wordLength + 1
			idxs[len(idxs)-1]++
		} else if length+wordLength == maxWidth {
			idxs[len(idxs)-1]++
			length = 0
			if i != len(words)-1 {
				idxs = append(idxs, 0)
			}
		} else {
			idxs = append(idxs, 1)
			length = wordLength + 1
		}
	}

	rows := make([][]string, len(idxs))
	row := 0
	wordIdx := 0 // word IDX in words list

	for ; row < len(idxs); row++ {
		// collect current "row" words
		currentWordsCnt, currWordsLength := idxs[row], 0
		currentWords := make([]string, currentWordsCnt)

		for i := range currentWordsCnt {
			currentWords[i] = words[wordIdx]
			currWordsLength += len(words[wordIdx])
			wordIdx++
		}

		// handle spaces in row
		spaces := maxWidth - currWordsLength

		// set spaces from first to pre-last
		buf := make([]int, currentWordsCnt)
		if row == len(idxs)-1 {
			for i := 0; i < currentWordsCnt-1; i++ {
				buf[i]++
				spaces--
			}
			buf[currentWordsCnt-1] = spaces // if it's last row left-justified
		} else {
			for i := 0; spaces > 0; spaces-- {
				buf[i]++
				i++
				if i >= currentWordsCnt-1 {
					i = 0
				}
			}
		}

		for i, word := range currentWords {
			if buf[i] > 0 {
				word = strings.Join([]string{word, strings.Repeat(" ", buf[i])}, "")
			}
			rows[row] = append(rows[row], word)
		}
	}

	res := make([]string, len(rows))
	for i, r := range rows {
		// fmt.Println(r)
		res[i] = strings.Join(r, "")
	}

	return res
}
