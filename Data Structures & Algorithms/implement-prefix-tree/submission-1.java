class PrefixTree {
    private PrefixTree[] children;
    private boolean isEnd;

    public PrefixTree() {
        children = new PrefixTree[26];
        isEnd = false;
    }

    public void insert(String word) {
        PrefixTree node = this;
        for (char c : word.toCharArray()) {
            int i = c - 'a';
            if (node.children[i] == null) {
                node.children[i] = new PrefixTree();
            }
            node = node.children[i];
        }
        node.isEnd = true;
    }

    public boolean search(String word) {
        PrefixTree node = find(word);
        return node != null && node.isEnd;
    }

    public boolean startsWith(String prefix) {
        return find(prefix) != null;
    }

    private PrefixTree find(String word) {
        PrefixTree node = this;
        for (char c : word.toCharArray()) {
            int i = c - 'a';
            if (node.children[i] == null) {
                return null;
            }
            node = node.children[i];
        }
        return node;
    }
}