class Solution {
    public int evalRPN(String[] tokens) {
        Deque<Integer> stack = new ArrayDeque<>();

        for (String tok : tokens) {
            switch (tok) {
                case "+": {
                    int b = stack.pop(), a = stack.pop();
                    stack.push(a + b);
                    break;
                }
                case "-": {
                    int b = stack.pop(), a = stack.pop();
                    stack.push(a - b);
                    break;
                }
                case "*": {
                    int b = stack.pop(), a = stack.pop();
                    stack.push(a * b);
                    break;
                }
                case "/": {
                    int b = stack.pop(), a = stack.pop();
                    stack.push(a / b); // Java division already truncates toward zero
                    break;
                }
                default:
                    stack.push(Integer.parseInt(tok));
            }
        }

        return stack.pop();
    }
}