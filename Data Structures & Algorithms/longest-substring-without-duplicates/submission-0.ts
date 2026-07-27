class Solution {
    lengthOfLongestSubstring(s: string): number {
        const window = new Set<string>();
        let l = 0;
        let res = 0;

        for (let r = 0; r < s.length; r++) {
            while (window.has(s[r])) {
                window.delete(s[l]);
                l++;
            }
            window.add(s[r]);
            res = Math.max(res, r - l + 1);
        }

        return res;
    }
}