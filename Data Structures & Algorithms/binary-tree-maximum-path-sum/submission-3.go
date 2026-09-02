func maxPathSum(root *TreeNode) int {
    best := math.MinInt32

    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := max(dfs(node.Left), 0)
        right := max(dfs(node.Right), 0)
        best = max(best, node.Val+left+right)
        return node.Val + max(left, right)
    }

    dfs(root)
    return best
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}