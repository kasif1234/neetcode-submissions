class Solution {
    maxArea(heights: number[]): number {
        let left = 0;
        let right = heights.length - 1;
        let maxWater = 0;

        while (left < right) {
            const width = right - left;
            const height = Math.min(heights[left], heights[right]);
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