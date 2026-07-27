func lengthOfLongestSubstring(s string) int {
    window := make(map[byte]bool)
    l, res := 0, 0

    for r := 0; r < len(s); r++ {
        for window[s[r]] {
            delete(window, s[l])
            l++
        }
        window[s[r]] = true
        if r-l+1 > res {
            res = r - l + 1
        }
    }

    return res
}