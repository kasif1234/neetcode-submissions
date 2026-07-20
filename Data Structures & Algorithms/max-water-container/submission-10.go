func maxArea(heights []int) int {
    left, right := 0, len(heights)-1
    maxWater := 0

    for left < right {
        width := right - left
        height := heights[left]
        if heights[right] < height {
            height = heights[right]
        }
        area := width * height
        if area > maxWater {
            maxWater = area
        }

        if heights[left] < heights[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}