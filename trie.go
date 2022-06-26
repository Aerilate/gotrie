package trie

type trie struct {
	endOfString bool
	count int
	children map[rune]*trie
}

func NewTrie() *trie {
	root := &trie{
		children: make(map[rune]*trie),
	}
	return root
}

func (t *trie) Insert(word string) {
	curr := t
	for _, char := range word {
		child, ok := curr.children[char]
		if !ok {
			curr.children[char] = NewTrie()
			child = curr.children[char]
		}
		child.count++
		curr = child
	}
	curr.endOfString = true
}

func (t *trie) Remove(word string) bool {
	if !t.Contains(word) {
		return false
	}
	curr := t
	for _, char := range word {
		child, ok := curr.children[char]
		if !ok {
			return false
		} else if child.count == 1 {
			delete(curr.children, char)
			return true
		}
		child.count--
		curr = child
	}
	curr.endOfString = false
	return true
}

func (t trie) Contains(word string) bool {
	curr := t
	for _, char := range word {
		child, ok := curr.children[char]
		if !ok {
			return false
		}
		curr = *child
	}
	return curr.endOfString
}

func (t trie) StartsWith(word string) bool {
	curr := t
	for _, char := range word {
		child, ok := curr.children[char]
		if !ok {
			return false
		}
		curr = *child
	}
	return true
}
