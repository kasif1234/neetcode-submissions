func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    
    l, r := 0, len(height)-1
    leftMax, rightMax := height[l], height[r]
    res := 0
    
    for l < r {
        if leftMax < rightMax {
            l++
            if height[l] > leftMax {
                leftMax = height[l]
            }
            res += leftMax - height[l]
        } else {
            r--
            if height[r] > rightMax {
                rightMax = height[r]
            }
            res += rightMax - height[r]
        }
    }
    
    return res
}