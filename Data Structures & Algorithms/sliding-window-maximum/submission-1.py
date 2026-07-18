from collections import deque

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        result = []
        dq = deque()  # stores indices, values in decreasing order

        for i, num in enumerate(nums):
            # Remove indices out of the current window from the front
            while dq and dq[0] < i - k + 1:
                dq.popleft()

            # Remove smaller values from the back
            while dq and nums[dq[-1]] < num:
                dq.pop()

            dq.append(i)

            if i >= k - 1:
                result.append(nums[dq[0]])

        return result