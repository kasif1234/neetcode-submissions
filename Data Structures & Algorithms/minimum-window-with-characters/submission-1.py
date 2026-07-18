class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not t or len(s) < len(t):
            return ""

        need = Counter(t)
        window = {}
        have = 0
        required = len(need)

        res_len = float("inf")
        res_left = 0

        left = 0
        for right in range(len(s)):
            c = s[right]
            window[c] = window.get(c, 0) + 1

            if c in need and window[c] == need[c]:
                have += 1

            while have == required:
                if right - left + 1 < res_len:
                    res_len = right - left + 1
                    res_left = left

                left_char = s[left]
                window[left_char] -= 1
                if left_char in need and window[left_char] < need[left_char]:
                    have -= 1
                left += 1

        return "" if res_len == float("inf") else s[res_left:res_left + res_len]