public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        buildString(root, sb);
        return sb.toString();
    }

    private void buildString(TreeNode node, StringBuilder sb) {
        if (node == null) {
            sb.append("N,");
        } else {
            sb.append(node.val).append(",");
            buildString(node.left, sb);
            buildString(node.right, sb);
        }
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        Queue<String> tokens = new LinkedList<>(Arrays.asList(data.split(",")));
        return buildTree(tokens);
    }

    private TreeNode buildTree(Queue<String> tokens) {
        String val = tokens.poll();
        if (val.equals("N")) {
            return null;
        }
        TreeNode node = new TreeNode(Integer.parseInt(val));
        node.left = buildTree(tokens);
        node.right = buildTree(tokens);
        return node;
    }
}