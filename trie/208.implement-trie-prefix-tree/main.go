package main

import "fmt"

/**
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings.
There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:
- Trie() Initializes the trie object.
- void insert(String word) Inserts the string word into the trie.
- boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
- boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


Example 1:
	Input
	["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
	[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
	Output
	[null, null, true, false, true, null, true]

	Explanation
	Trie trie = new Trie();
	trie.insert("apple");
	trie.search("apple");   // return True
	trie.search("app");     // return False
	trie.startsWith("app"); // return True
	trie.insert("app");
	trie.search("app");     // return True


Constraints:
	1 <= word.length, prefix.length <= 2000
	word and prefix consist only of lowercase English letters.
	At most 3 * 104 calls in total will be made to insert, search, and startsWith.
*/

func main() {
	fmt.Println(Constructor())
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		i := int(c - 'a')
		if curr.children[i] == nil {
			curr.children[i] = &Trie{}
		}
		curr = curr.children[i]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this.find(word)
	return curr != nil && curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) != nil
}

func (this *Trie) find(s string) *Trie {
	curr := this
	for _, c := range s {
		i := int(c - 'a')
		if curr.children[i] == nil {
			return nil
		}
		curr = curr.children[i]
	}
	return curr
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
