func buildTree(preorder []int, inorder []int) *TreeNode {
    inMap := make(map[int]int)
    for i, v := range inorder {
        inMap[v] = i
    }
    preIdx := 0

    var build func(l, r int) *TreeNode
    build = func(l, r int) *TreeNode {
        if l > r {
            return nil
        }

        rootVal := preorder[preIdx]
        preIdx++
        root := &TreeNode{Val: rootVal}

        mid := inMap[rootVal]
        root.Left = build(l, mid-1)
        root.Right = build(mid+1, r)

        return root
    }

    return build(0, len(inorder)-1)
}