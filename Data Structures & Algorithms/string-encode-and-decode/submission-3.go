import (
    "strconv"
    "strings"
)

type Solution struct{}

func (Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, s := range strs {
        sb.WriteString(strconv.Itoa(len(s)))
        sb.WriteByte('#')
        sb.WriteString(s)
    }
    return sb.String()
}

func (Solution) Decode(s string) []string {
    res := []string{}
    i := 0
    for i < len(s) {
        j := i
        for s[j] != '#' {
            j++
        }
        length, _ := strconv.Atoi(s[i:j])
        i = j + 1
        res = append(res, s[i:i+length])
        i += length
    }
    return res
}