class TrieNode {
    children: Map<string, TrieNode> = new Map();
    word: string | null = null;
}

class Solution {
    findWords(board: string[][], words: string[]): string[] {
        const root = new TrieNode();
        for (const w of words) {
            let node = root;
            for (const ch of w) {
                if (!node.children.has(ch)) node.children.set(ch, new TrieNode());
                node = node.children.get(ch)!;
            }
            node.word = w;
        }

        const ROWS = board.length, COLS = board[0].length;
        const res: string[] = [];

        const dfs = (r: number, c: number, node: TrieNode): void => {
            const ch = board[r][c];
            const child = node.children.get(ch);
            if (!child) return;

            if (child.word !== null) {
                res.push(child.word);
                child.word = null;
            }

            board[r][c] = '#';
            for (const [dr, dc] of [[1,0],[-1,0],[0,1],[0,-1]]) {
                const nr = r + dr, nc = c + dc;
                if (nr >= 0 && nr < ROWS && nc >= 0 && nc < COLS && board[nr][nc] !== '#') {
                    dfs(nr, nc, child);
                }
            }
            board[r][c] = ch;

            if (child.children.size === 0) node.children.delete(ch);
        };

        for (let r = 0; r < ROWS; r++)
            for (let c = 0; c < COLS; c++)
                dfs(r, c, root);

        return res;
    }
}