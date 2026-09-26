type TrieNode struct {
    children [26]*TrieNode
    word     string
}

func findWords(board [][]byte, words []string) []string {
    root := &TrieNode{}
    for _, w := range words {
        node := root
        for i := 0; i < len(w); i++ {
            idx := w[i] - 'a'
            if node.children[idx] == nil {
                node.children[idx] = &TrieNode{}
            }
            node = node.children[idx]
        }
        node.word = w
    }

    ROWS, COLS := len(board), len(board[0])
    res := []string{}

    var dfs func(r, c int, node *TrieNode)
    dfs = func(r, c int, node *TrieNode) {
        idx := board[r][c] - 'a'
        if node.children[idx] == nil {
            return
        }
        child := node.children[idx]
        if child.word != "" {
            res = append(res, child.word)
            child.word = ""
        }

        ch := board[r][c]
        board[r][c] = '#'
        dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < ROWS && nc >= 0 && nc < COLS && board[nr][nc] != '#' {
                dfs(nr, nc, child)
            }
        }
        board[r][c] = ch

        hasChild := false
        for i := 0; i < 26; i++ {
            if child.children[i] != nil {
                hasChild = true
                break
            }
        }
        if !hasChild {
            node.children[idx] = nil
        }
    }

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            dfs(r, c, root)
        }
    }

    return res
}