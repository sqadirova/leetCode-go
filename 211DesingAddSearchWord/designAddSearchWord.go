// Design Add and Search Words Data Structure
package main

import "fmt"

type WordDictionaryNode struct {
	children    map[rune]*WordDictionaryNode
	isEndOfWord bool
}

type WordDictionary struct {
	root *WordDictionaryNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &WordDictionaryNode{
			children: make(map[rune]*WordDictionaryNode),
		},
	}

}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, char := range word {
		if _, ok := curr.children[char]; !ok {
			curr.children[char] = &WordDictionaryNode{
				children: make(map[rune]*WordDictionaryNode),
			}
		}
		curr = curr.children[char]
	}

	curr.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchInNode(word, this.root)
}

func (this *WordDictionary) SearchInNode(word string, node *WordDictionaryNode) bool {
	if len(word) == 0 {
		return node.isEndOfWord
	}

	char := rune(word[0])
	remainingWord := word[1:]

	if char == '.' {
		for _, childNode := range node.children {
			if this.SearchInNode(remainingWord, childNode) {
				return true
			}
		}
		return false
	} else {
		if childNode, ok := node.children[char]; ok {
			return this.SearchInNode(remainingWord, childNode)
		}
		return false
	}

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	wordDictionary := Constructor()

	// Example from the problem description
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")

	fmt.Printf("Search 'pad': %t (Expected: false)\n", wordDictionary.Search("pad")) // return False
	fmt.Printf("Search 'bad': %t (Expected: true)\n", wordDictionary.Search("bad"))  // return True
	fmt.Printf("Search '.ad': %t (Expected: true)\n", wordDictionary.Search(".ad"))  // return True
	fmt.Printf("Search 'b..': %t (Expected: true)\n", wordDictionary.Search("b.."))  // return True

	fmt.Println("\n--- Additional Tests ---")
	wordDictionary.AddWord("apple")
	wordDictionary.AddWord("apply")
	fmt.Printf("Search 'apple': %t (Expected: true)\n", wordDictionary.Search("apple"))
	fmt.Printf("Search 'app.y': %t (Expected: true)\n", wordDictionary.Search("app.y"))
	fmt.Printf("Search 'ap.le': %t (Expected: true)\n", wordDictionary.Search("ap.le"))
	fmt.Printf("Search 'ap.ly': %t (Expected: true)\n", wordDictionary.Search("ap.ly"))
	fmt.Printf("Search 'a.p.e': %t (Expected: true)\n", wordDictionary.Search("a.p.e"))
	fmt.Printf("Search 'a..y': %t (Expected: true)\n", wordDictionary.Search("a..y"))
	fmt.Printf("Search 'appl.': %t (Expected: true)\n", wordDictionary.Search("appl."))
	fmt.Printf("Search 'applee': %t (Expected: false)\n", wordDictionary.Search("applee"))
	fmt.Printf("Search 'app': %t (Expected: false)\n", wordDictionary.Search("app"))
	fmt.Printf("Search 'a..': %t (Expected: true)\n", wordDictionary.Search("a.."))     // Should match "app", "bad", "dad", "mad"
	fmt.Printf("Search '...': %t (Expected: true)\n", wordDictionary.Search("..."))     // Should match "bad", "dad", "mad"
	fmt.Printf("Search '.....': %t (Expected: true)\n", wordDictionary.Search(".....")) // Should match "apple", "apply"
	fmt.Printf("Search '......': %t (Expected: false)\n", wordDictionary.Search("......"))
}
