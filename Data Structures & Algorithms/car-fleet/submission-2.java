class Solution {
    public int carFleet(int target, int[] position, int[] speed) {
        int n = position.length;
        Integer[] indices = new Integer[n];
        for (int i = 0; i < n; i++) indices[i] = i;
        
        Arrays.sort(indices, (a, b) -> position[b] - position[a]);
        
        Deque<Double> stack = new ArrayDeque<>();
        
        for (int i : indices) {
            double time = (double) (target - position[i]) / speed[i];
            if (stack.isEmpty() || time > stack.peek()) {
                stack.push(time);
            }
        }
        
        return stack.size();
    }
}