type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{root: &TrieNode{}}
}

func (wd *WordDictionary) AddWord(word string) {
    cur := wd.root
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if cur.children[idx] == nil {
            cur.children[idx] = &TrieNode{}
        }
        cur = cur.children[idx]
    }
    cur.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
    var dfs func(node *TrieNode, i int) bool
    dfs = func(node *TrieNode, i int) bool {
        if node == nil {
            return false
        }
        if i == len(word) {
            return node.isEnd
        }
        c := word[i]
        if c == '.' {
            for _, child := range node.children {
                if child != nil && dfs(child, i+1) {
                    return true
                }
            }
            return false
        }
        return dfs(node.children[c-'a'], i+1)
    }
    return dfs(wd.root, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */