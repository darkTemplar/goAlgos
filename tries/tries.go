package main

import (
	"fmt"
)

type Trie struct {
	Key      rune
	Children map[rune]*Trie
}

func main() {
	// Uncomment below to test normal trie
	/*fmt.Println("Starting trie construction")
	trie := &Trie{Children: make(map[rune]*Trie)}
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.insert(words[i])
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		_, found := trie.search(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}*/
	fmt.Println("Starting suffix trie construction")
	suffixTrie := &Trie{Children: make(map[rune]*Trie)}
	suffixTrie.insertSuffix("banana")
	// test for suffix
	fmt.Printf("Checking if %s is a substring: %v\n", "ana", suffixTrie.hasSuffix("ana"))
	fmt.Printf("Checking if %s is a substring: %v\n", "ban", suffixTrie.hasSuffix("ban"))
}

func (trie *Trie) insert(key string) {
	fmt.Println("Inserting key:", key)
	if len(key) <= 0 {
		return
	}
	current := trie
	for _, r := range key {
		if child, ok := current.Children[r]; ok {
			current = child
		} else {
			newChild := &Trie{Key: r, Children: make(map[rune]*Trie)}
			current.Children[r] = newChild
			current = newChild
		}
	}
}

func (trie *Trie) search(key string) (*Trie, bool) {
	if len(key) <= 0 {
		return nil, true
	}
	current := trie
	for _, r := range key {
		if child, ok := current.Children[r]; ok {
			current = child
		} else {
			return current, false
		}
	}
	return current, true
}

func (trie *Trie) insertSuffix(key string) {
	// add special char at end
	key += "$"
	for i := 0; i < len(key); i++ {
		trie.insert(key[i:])
	}
}

func (trie *Trie) hasSuffix(key string) bool {
	if node, found := trie.search(key); found {
		_, ok := node.Children[rune('$')]
		return ok
	}
	return false
}

//TODO: contacts
//TODO: simple regex
