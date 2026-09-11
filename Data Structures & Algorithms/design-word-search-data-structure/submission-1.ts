class TrieNode {
    children: Map<string, TrieNode> = new Map();
    isEnd: boolean = false;
}

class WordDictionary {
    root: TrieNode;

    constructor() {
        this.root = new TrieNode();
    }

    /**
     * @param {string} word
     * @return {void}
     */
    addWord(word: string): void {
        let cur = this.root;
        for (const c of word) {
            if (!cur.children.has(c)) {
                cur.children.set(c, new TrieNode());
            }
            cur = cur.children.get(c)!;
        }
        cur.isEnd = true;
    }

    /**
     * @param {string} word
     * @return {boolean}
     */
    search(word: string): boolean {
        const dfs = (node: TrieNode, i: number): boolean => {
            if (i === word.length) {
                return node.isEnd;
            }
            const c = word[i];
            if (c === '.') {
                for (const child of node.children.values()) {
                    if (dfs(child, i + 1)) {
                        return true;
                    }
                }
                return false;
            } else {
                const next = node.children.get(c);
                if (!next) {
                    return false;
                }
                return dfs(next, i + 1);
            }
        };

        return dfs(this.root, 0);
    }
}