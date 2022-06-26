package trie

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func TestInsertBasic(t *testing.T) {
	trie := NewTrie()
	assert.False(t, trie.Contains("hello"))
	trie.Insert("hello")
	assert.True(t, trie.Contains("hello"))
}

func TestRemoveBasic(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	assert.True(t, trie.Contains("hello"))
	assert.True(t, trie.Remove("hello"))
	assert.False(t, trie.Contains("hello"))
}

func TestRemoveFail(t *testing.T) {
	trie := NewTrie()
	assert.False(t, trie.Remove("hello"))
}

func TestInsertSubstringBefore(t *testing.T) {
	trie := NewTrie()
	trie.Insert("he")
	trie.Insert("hello")
	assert.True(t, trie.Contains("he"))
	assert.True(t, trie.Contains("hello"))
}


func TestInsertSubstringAfter(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	assert.False(t, trie.Contains("he"))
	trie.Insert("he")
	assert.True(t, trie.Contains("he"))
	assert.True(t, trie.Contains("hello"))
}

func TestRemoveSubstringBefore(t *testing.T) {
	trie := NewTrie()
	trie.Insert("he")
	trie.Insert("hello")
	trie.Remove("he")
	assert.False(t, trie.Contains("he"))
	assert.True(t, trie.Contains("hello"))
}

func TestRemoveSubstringAfter(t *testing.T) {
	trie := NewTrie()
	trie.Insert("he")
	trie.Insert("hello")
	trie.Remove("hello")
	assert.True(t, trie.Contains("he"))
	assert.False(t, trie.Contains("hello"))
}

func TestStartsWith(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	assert.True(t, trie.StartsWith("he"))
	trie.Remove("hello")
	assert.False(t, trie.StartsWith("he"))
}

func TestStartsWithExact(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	assert.True(t, trie.StartsWith("hello"))
	trie.Remove("hello")
	assert.False(t, trie.StartsWith("hello"))
}
