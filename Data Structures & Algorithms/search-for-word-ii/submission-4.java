class TrieNode {
    TrieNode[] children = new TrieNode[26];
    String word = null;
}

class Solution {
    List<String> res = new ArrayList<>();
    int ROWS, COLS;

    public List<String> findWords(char[][] board, String[] words) {
        TrieNode root = new TrieNode();
        for (String w : words) {
            TrieNode node = root;
            for (char ch : w.toCharArray()) {
                int idx = ch - 'a';
                if (node.children[idx] == null) node.children[idx] = new TrieNode();
                node = node.children[idx];
            }
            node.word = w;
        }

        ROWS = board.length;
        COLS = board[0].length;

        for (int r = 0; r < ROWS; r++)
            for (int c = 0; c < COLS; c++)
                dfs(board, r, c, root);

        return res;
    }

    private void dfs(char[][] board, int r, int c, TrieNode node) {
        int idx = board[r][c] - 'a';
        if (node.children[idx] == null) return;
        TrieNode child = node.children[idx];

        if (child.word != null) {
            res.add(child.word);
            child.word = null;
        }

        char ch = board[r][c];
        board[r][c] = '#';
        int[][] dirs = {{1,0},{-1,0},{0,1},{0,-1}};
        for (int[] d : dirs) {
            int nr = r + d[0], nc = c + d[1];
            if (nr >= 0 && nr < ROWS && nc >= 0 && nc < COLS && board[nr][nc] != '#')
                dfs(board, nr, nc, child);
        }
        board[r][c] = ch;

        boolean hasChild = false;
        for (int i = 0; i < 26; i++) if (child.children[i] != null) { hasChild = true; break; }
        if (!hasChild) node.children[idx] = null;
    }
}