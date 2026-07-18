class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        max_area = 0
        stack = []  # stores indices, heights increasing

        for i in range(len(heights) + 1):
            curr_height = heights[i] if i < len(heights) else 0

            while stack and heights[stack[-1]] > curr_height:
                height = heights[stack.pop()]
                width = i if not stack else i - stack[-1] - 1
                max_area = max(max_area, height * width)

            stack.append(i)

        return max_area