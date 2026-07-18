class Solution {
public:
    string minWindow(string s, string t) {
        if (t.empty() || s.length() < t.length()) return "";

        unordered_map<char, int> need;
        for (char c : t) {
            need[c]++;
        }

        unordered_map<char, int> window;
        int have = 0;
        int required = need.size();

        int resLen = INT_MAX;
        int resLeft = 0;

        int left = 0;
        for (int right = 0; right < s.length(); right++) {
            char c = s[right];
            window[c]++;

            if (need.count(c) && window[c] == need[c]) {
                have++;
            }

            while (have == required) {
                if (right - left + 1 < resLen) {
                    resLen = right - left + 1;
                    resLeft = left;
                }

                char leftChar = s[left];
                window[leftChar]--;
                if (need.count(leftChar) && window[leftChar] < need[leftChar]) {
                    have--;
                }
                left++;
            }
        }

        return resLen == INT_MAX ? "" : s.substr(resLeft, resLen);
    }
};