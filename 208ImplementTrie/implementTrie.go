// Implement Trie (Prefix Tree)
package main

import "fmt"

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, char := range word {
		if _, ok := curr.children[char]; !ok {
			curr.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		curr = curr.children[char]
	}
	curr.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, char := range word {
		if _, ok := curr.children[char]; !ok {
			return false
		}
		curr = curr.children[char]
	}
	return curr.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, char := range prefix {
		if _, ok := curr.children[char]; !ok {
			return false
		}
		curr = curr.children[char]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()

	// Example 1 from the problem description
	trie.Insert("apple")
	fmt.Printf("Search 'apple': %t (Expected: true)\n", trie.Search("apple"))     // return True
	fmt.Printf("Search 'app': %t (Expected: false)\n", trie.Search("app"))        // return False
	fmt.Printf("StartsWith 'app': %t (Expected: true)\n", trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Printf("Search 'app': %t (Expected: true)\n", trie.Search("app")) // return True

	fmt.Println("\n--- Additional Tests ---")
	trie.Insert("banana")
	fmt.Printf("Search 'banana': %t (Expected: true)\n", trie.Search("banana"))
	fmt.Printf("Search 'ban': %t (Expected: false)\n", trie.Search("ban"))
	fmt.Printf("StartsWith 'ban': %t (Expected: true)\n", trie.StartsWith("ban"))
	fmt.Printf("StartsWith 'band': %t (Expected: false)\n", trie.StartsWith("band"))
	fmt.Printf("Search 'ap': %t (Expected: false)\n", trie.Search("ap"))
	fmt.Printf("StartsWith 'ap': %t (Expected: true)\n", trie.StartsWith("ap"))
}
