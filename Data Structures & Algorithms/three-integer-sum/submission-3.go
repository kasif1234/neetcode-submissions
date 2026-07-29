func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    n := len(nums)

    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        j, k := i+1, n-1
        for j < k {
            total := nums[i] + nums[j] + nums[k]
            if total < 0 {
                j++
            } else if total > 0 {
                k--
            } else {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j] == nums[j-1] {
                    j++
                }
                for j < k && nums[k] == nums[k+1] {
                    k--
                }
            }
        }
    }
    return res
}