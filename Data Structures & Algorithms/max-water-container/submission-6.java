class Solution {
    public int maxArea(int[] heights) {
        int left = 0;
        int right = heights.length - 1;
        int maxWater = 0;

        while (left < right) {
            int width = right - left;
            int height = Math.min(heights[left], heights[right]);
            maxWater = Math.max(maxWater, width * height);

            if (heights[left] < heights[right]) {
                left++;
            } else {
                right--;
            }
        }

        return maxWater;
    }
}