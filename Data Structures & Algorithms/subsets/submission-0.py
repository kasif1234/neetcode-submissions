class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        subset = []

        def dfs(i):
            if i == len(nums):
                res.append(subset.copy())  # save a copy
                return

            # choice 1: take nums[i]
            subset.append(nums[i])
            dfs(i + 1)

            # undo, then choice 2: skip nums[i]
            subset.pop()
            dfs(i + 1)

        dfs(0)
        return res