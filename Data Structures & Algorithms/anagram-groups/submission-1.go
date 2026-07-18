func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)

    for _, s := range strs {
        var count [26]int
        for _, c := range s {
            count[c-'a']++
        }
        m[count] = append(m[count], s)
    }

    result := make([][]string, 0, len(m))
    for _, group := range m {
        result = append(result, group)
    }
    return result
}