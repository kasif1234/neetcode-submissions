func leastInterval(tasks []byte, n int) int {
    count := make([]int, 26)
    for _, t := range tasks {
        count[t-'A']++
    }

    maxFreq := 0
    for _, c := range count {
        if c > maxFreq {
            maxFreq = c
        }
    }

    maxCount := 0
    for _, c := range count {
        if c == maxFreq {
            maxCount++
        }
    }

    res := (maxFreq-1)*(n+1) + maxCount
    if len(tasks) > res {
        return len(tasks)
    }
    return res
}