type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("N,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		val := vals[idx]
		idx++
		if val == "N" {
			return nil
		}
		num, _ := strconv.Atoi(val)
		node := &TreeNode{Val: num}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	return dfs()
}