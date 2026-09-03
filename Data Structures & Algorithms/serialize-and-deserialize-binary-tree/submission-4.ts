class Codec {
    /**
     * Encodes a tree to a single string.
     */
    serialize(root: TreeNode | null): string {
        const res: string[] = [];

        const dfs = (node: TreeNode | null): void => {
            if (node === null) {
                res.push("N");
                return;
            }
            res.push(node.val.toString());
            dfs(node.left);
            dfs(node.right);
        };

        dfs(root);
        return res.join(",");
    }

    /**
     * Decodes your encoded data to tree.
     */
    deserialize(data: string): TreeNode | null {
        const vals = data.split(",");
        let idx = 0;

        const dfs = (): TreeNode | null => {
            const val = vals[idx];
            idx++;
            if (val === "N") {
                return null;
            }
            const node = new TreeNode(parseInt(val));
            node.left = dfs();
            node.right = dfs();
            return node;
        };

        return dfs();
    }
}