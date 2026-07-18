func maxSlidingWindow(nums []int, k int) []int {
    var result []int
    var deque []int // stores indices, values in decreasing order

    for i := 0; i < len(nums); i++ {
        // Remove indices out of the current window from the front
        for len(deque) > 0 && deque[0] < i-k+1 {
            deque = deque[1:]
        }

        // Remove smaller values from the back
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
            deque = deque[:len(deque)-1]
        }

        deque = append(deque, i)

        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }

    return result
}