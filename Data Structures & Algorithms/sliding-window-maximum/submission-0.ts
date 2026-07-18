class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    maxSlidingWindow(nums: number[], k: number): number[] {
        const result: number[] = [];
        const deque: number[] = []; // stores indices, values in decreasing order

        for (let i = 0; i < nums.length; i++) {
            // Remove indices out of the current window from the front
            while (deque.length && deque[0] < i - k + 1) {
                deque.shift();
            }

            // Remove smaller values from the back since they can never be the max
            while (deque.length && nums[deque[deque.length - 1]] < nums[i]) {
                deque.pop();
            }

            deque.push(i);

            if (i >= k - 1) {
                result.push(nums[deque[0]]);
            }
        }

        return result;
    }
}