package trie

import "testing"

func TestTrie_InsertAndSearch(t *testing.T) {
	trie := NewTrie()

	wordsToInsert := []string{"hello", "helium", "hero"}
	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	testCases := []struct {
		word string
		want bool
	}{
		{"hello", true},
		{"helium", true},
		{"hero", true},
		{"her", false},
		{"he", false},
		{"world", false},
	}

	for _, tc := range testCases {
		got := trie.Search(tc.word)
		if got != tc.want {
			t.Errorf("Search(%q) = %v; want %v", tc.word, got, tc.want)
		}
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie := NewTrie()

	prefixesToInsert := []string{"hello", "helium", "hero"}
	for _, prefix := range prefixesToInsert {
		trie.Insert(prefix)
	}

	testCases := []struct {
		prefix string
		want   bool
	}{
		{"he", true},
		{"hell", true},
		{"hero", true},
		{"herox", false},
		{"hi", false},
	}

	for _, tc := range testCases {
		got := trie.StartsWith(tc.prefix)
		if got != tc.want {
			t.Errorf("StartsWith(%q) = %v; want %v", tc.prefix, got, tc.want)
		}
	}
}
