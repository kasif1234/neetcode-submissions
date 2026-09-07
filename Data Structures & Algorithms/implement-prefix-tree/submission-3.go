type PrefixTree struct {
    children [26]*PrefixTree
    isEnd    bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (t *PrefixTree) Insert(word string) {
    node := t
    for _, c := range word {
        i := c - 'a'
        if node.children[i] == nil {
            node.children[i] = &PrefixTree{}
        }
        node = node.children[i]
    }
    node.isEnd = true
}

func (t *PrefixTree) find(word string) *PrefixTree {
    node := t
    for _, c := range word {
        i := c - 'a'
        if node.children[i] == nil {
            return nil
        }
        node = node.children[i]
    }
    return node
}

func (t *PrefixTree) Search(word string) bool {
    node := t.find(word)
    return node != nil && node.isEnd
}

func (t *PrefixTree) StartsWith(prefix string) bool {
    return t.find(prefix) != nil
}