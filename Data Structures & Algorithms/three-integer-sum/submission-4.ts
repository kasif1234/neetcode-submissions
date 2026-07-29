class Solution {
    threeSum(nums: number[]): number[][] {
        nums.sort((a, b) => a - b);
        const res: number[][] = [];
        const n = nums.length;

        for (let i = 0; i < n; i++) {
            if (nums[i] > 0) break;
            if (i > 0 && nums[i] === nums[i - 1]) continue;

            let j = i + 1, k = n - 1;
            while (j < k) {
                const total = nums[i] + nums[j] + nums[k];
                if (total < 0) {
                    j++;
                } else if (total > 0) {
                    k--;
                } else {
                    res.push([nums[i], nums[j], nums[k]]);
                    j++;
                    k--;
                    while (j < k && nums[j] === nums[j - 1]) j++;
                    while (j < k && nums[k] === nums[k + 1]) k--;
                }
            }
        }
        return res;
    }
}