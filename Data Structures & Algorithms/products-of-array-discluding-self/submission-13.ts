/**
 * @param {number[]} nums
 * @return {number[]}
 */
class Solution {
    productExceptSelf(nums: number[]): number[] {
        const n = nums.length;
        const output = new Array(n).fill(1);

        for (let i = 1; i < n; i++) {
            output[i] = output[i - 1] * nums[i - 1];
        }

        let suffix = 1;
        for (let i = n - 1; i >= 0; i--) {
            output[i] *= suffix;
            suffix *= nums[i];
        }

        return output;
    }
}