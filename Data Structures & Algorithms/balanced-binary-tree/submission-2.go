/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return dfs(root) != -1
}

func dfs(node *TreeNode) int {
    if node == nil {
        return 0
    }

    left := dfs(node.Left)
    if left == -1 {
        return -1
    }

    right := dfs(node.Right)
    if right == -1 {
        return -1
    }

    diff := left - right
    if diff > 1 || diff < -1 {
        return -1
    }

    if left > right {
        return left + 1
    }
    return right + 1
}