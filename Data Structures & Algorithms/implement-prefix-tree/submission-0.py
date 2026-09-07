class PrefixTree:

    def __init__(self):
        self.children = [None] * 26
        self.isEnd = False

    def insert(self, word: str) -> None:
        node = self
        for c in word:
            i = ord(c) - ord('a')
            if node.children[i] is None:
                node.children[i] = PrefixTree()
            node = node.children[i]
        node.isEnd = True

    def search(self, word: str) -> bool:
        node = self._find(word)
        return node is not None and node.isEnd

    def startsWith(self, prefix: str) -> bool:
        return self._find(prefix) is not None

    def _find(self, word: str):
        node = self
        for c in word:
            i = ord(c) - ord('a')
            if node.children[i] is None:
                return None
            node = node.children[i]
        return node