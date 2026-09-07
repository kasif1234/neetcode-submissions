class PrefixTree {
private:
    vector<PrefixTree*> children;
    bool isEnd;

    PrefixTree* find(const string& word) {
        PrefixTree* node = this;
        for (char c : word) {
            int i = c - 'a';
            if (node->children[i] == nullptr) {
                return nullptr;
            }
            node = node->children[i];
        }
        return node;
    }

public:
    PrefixTree() : children(26, nullptr), isEnd(false) {}

    void insert(string word) {
        PrefixTree* node = this;
        for (char c : word) {
            int i = c - 'a';
            if (node->children[i] == nullptr) {
                node->children[i] = new PrefixTree();
            }
            node = node->children[i];
        }
        node->isEnd = true;
    }

    bool search(string word) {
        PrefixTree* node = find(word);
        return node != nullptr && node->isEnd;
    }

    bool startsWith(string prefix) {
        return find(prefix) != nullptr;
    }
};