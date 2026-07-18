func minWindow(s string, t string) string {
    if len(t) == 0 || len(s) < len(t) {
        return ""
    }

    need := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        need[t[i]]++
    }

    window := make(map[byte]int)
    have := 0
    required := len(need)

    resLen := math.MaxInt32
    resLeft := 0

    left := 0
    for right := 0; right < len(s); right++ {
        c := s[right]
        window[c]++

        if needCount, ok := need[c]; ok && window[c] == needCount {
            have++
        }

        for have == required {
            if right-left+1 < resLen {
                resLen = right - left + 1
                resLeft = left
            }

            leftChar := s[left]
            window[leftChar]--
            if needCount, ok := need[leftChar]; ok && window[leftChar] < needCount {
                have--
            }
            left++
        }
    }

    if resLen == math.MaxInt32 {
        return ""
    }
    return s[resLeft : resLeft+resLen]
}