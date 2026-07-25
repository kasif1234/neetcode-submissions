func characterReplacement(s string, k int) int {
    var count [26]int
    maxFreq := 0
    left := 0
    maxLen := 0

    for right := 0; right < len(s); right++ {
        idx := s[right] - 'A'
        count[idx]++
        if count[idx] > maxFreq {
            maxFreq = count[idx]
        }

        for (right-left+1)-maxFreq > k {
            count[s[left]-'A']--
            left++
        }

        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }

    return maxLen
}