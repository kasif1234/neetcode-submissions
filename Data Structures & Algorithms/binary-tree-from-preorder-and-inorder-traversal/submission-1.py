class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        inMap = {val: idx for idx, val in enumerate(inorder)}
        self.preIdx = 0

        def build(l, r):
            if l > r:
                return None

            rootVal = preorder[self.preIdx]
            self.preIdx += 1
            root = TreeNode(rootVal)

            mid = inMap[rootVal]
            root.left = build(l, mid - 1)
            root.right = build(mid + 1, r)

            return root

        return build(0, len(inorder) - 1)