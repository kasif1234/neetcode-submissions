class Solution {
public:
    int characterReplacement(string s, int k) {
        vector<int> count(26, 0);
        int maxFreq = 0;
        int left = 0;
        int maxLen = 0;

        for (int right = 0; right < (int)s.size(); right++) {
            int idx = s[right] - 'A';
            count[idx]++;
            maxFreq = max(maxFreq, count[idx]);

            while ((right - left + 1) - maxFreq > k) {
                count[s[left] - 'A']--;
                left++;
            }

            maxLen = max(maxLen, right - left + 1);
        }

        return maxLen;
    }
};