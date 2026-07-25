class Solution {
    characterReplacement(s: string, k: number): number {
        const count = new Array(26).fill(0);
        let maxFreq = 0;
        let left = 0;
        let maxLen = 0;

        for (let right = 0; right < s.length; right++) {
            const idx = s.charCodeAt(right) - 65; // 'A' = 65
            count[idx]++;
            maxFreq = Math.max(maxFreq, count[idx]);

            // window size - maxFreq = number of chars that need replacing
            while ((right - left + 1) - maxFreq > k) {
                count[s.charCodeAt(left) - 65]--;
                left++;
            }

            maxLen = Math.max(maxLen, right - left + 1);
        }

        return maxLen;
    }
}