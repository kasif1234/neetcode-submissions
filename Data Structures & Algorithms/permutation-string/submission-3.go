func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var s1Count, s2Count [26]int

    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }

    matches := 0
    for i := 0; i < 26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }

    l := 0
    for r := len(s1); r < len(s2); r++ {
        if matches == 26 {
            return true
        }

        indexR := s2[r] - 'a'
        s2Count[indexR]++
        if s2Count[indexR] == s1Count[indexR] {
            matches++
        } else if s2Count[indexR] == s1Count[indexR]+1 {
            matches--
        }

        indexL := s2[l] - 'a'
        s2Count[indexL]--
        if s2Count[indexL] == s1Count[indexL] {
            matches++
        } else if s2Count[indexL] == s1Count[indexL]-1 {
            matches--
        }

        l++
    }

    return matches == 26
}