func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    
    isAlnum := func(c byte) bool {
        return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
    }
    
    toLower := func(c byte) byte {
        if c >= 'A' && c <= 'Z' {
            return c + ('a' - 'A')
        }
        return c
    }
    
    for left < right {
        for left < right && !isAlnum(s[left]) {
            left++
        }
        for left < right && !isAlnum(s[right]) {
            right--
        }
        
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }
        
        left++
        right--
    }
    
    return true
}