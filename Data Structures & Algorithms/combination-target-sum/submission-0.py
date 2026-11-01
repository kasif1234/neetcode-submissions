class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        res = []

        def dfs(start, cur, total):
            if total == target:
                res.append(cur[:])      # save a copy
                return
            for j in range(start, len(nums)):
                if total + nums[j] > target:
                    continue            # would overshoot, skip
                cur.append(nums[j])
                dfs(j, cur, total + nums[j])   # j, not j+1: reuse allowed
                cur.pop()               # undo the choice (backtrack)

        dfs(0, [], 0)
        return res