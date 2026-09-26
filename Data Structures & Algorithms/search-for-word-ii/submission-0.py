class TrieNode:
    def __init__(self):
        self.children = {}
        self.word = None

class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        root = TrieNode()
        for w in words:
            node = root
            for ch in w:
                node = node.children.setdefault(ch, TrieNode())
            node.word = w

        ROWS, COLS = len(board), len(board[0])
        res = []

        def dfs(r, c, node):
            ch = board[r][c]
            child = node.children.get(ch)
            if not child:
                return
            if child.word is not None:
                res.append(child.word)
                child.word = None  # avoid duplicate adds

            board[r][c] = '#'
            for dr, dc in ((1,0),(-1,0),(0,1),(0,-1)):
                nr, nc = r+dr, c+dc
                if 0 <= nr < ROWS and 0 <= nc < COLS and board[nr][nc] != '#':
                    dfs(nr, nc, child)
            board[r][c] = ch

            if not child.children:  # prune dead leaf
                del node.children[ch]

        for r in range(ROWS):
            for c in range(COLS):
                dfs(r, c, root)

        return res