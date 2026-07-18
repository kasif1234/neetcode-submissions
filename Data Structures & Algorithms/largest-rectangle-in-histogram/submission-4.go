func largestRectangleArea(heights []int) int {
    maxArea := 0
    stack := []int{} // stores indices, heights increasing

    for i := 0; i <= len(heights); i++ {
        currHeight := 0
        if i < len(heights) {
            currHeight = heights[i]
        }

        for len(stack) > 0 && heights[stack[len(stack)-1]] > currHeight {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]

            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            area := height * width
            if area > maxArea {
                maxArea = area
            }
        }

        stack = append(stack, i)
    }

    return maxArea
}