class Solution {
    search(nums: number[], target: number): number {
        let l = 0, r = nums.length - 1;

        while (l <= r) {
            const mid = Math.floor(l + (r - l) / 2);
            if (nums[mid] === target) return mid;

            // Left half is sorted
            if (nums[l] <= nums[mid]) {
                if (nums[l] <= target && target < nums[mid]) {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            } else {
                // Right half is sorted
                if (nums[mid] < target && target <= nums[r]) {
                    l = mid + 1;
                } else {
                    r = mid - 1;
                }
            }
        }

        return -1;
    }
}