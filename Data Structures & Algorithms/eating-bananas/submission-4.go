func minEatingSpeed(piles []int, h int) int {
    left, right := 1, 0
    for _, pile := range piles {
        if pile > right {
            right = pile
        }
    }
    res := right

    for left <= right {
        k := left + (right-left)/2
        hours := 0
        for _, pile := range piles {
            hours += (pile + k - 1) / k // ceil division
        }

        if hours <= h {
            res = k
            right = k - 1
        } else {
            left = k + 1
        }
    }

    return res
}