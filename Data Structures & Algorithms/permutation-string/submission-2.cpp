class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        if (s1.size() > s2.size()) return false;

        vector<int> s1Count(26, 0), s2Count(26, 0);

        for (int i = 0; i < s1.size(); i++) {
            s1Count[s1[i] - 'a']++;
            s2Count[s2[i] - 'a']++;
        }

        int matches = 0;
        for (int i = 0; i < 26; i++) {
            if (s1Count[i] == s2Count[i]) matches++;
        }

        int l = 0;
        for (int r = s1.size(); r < s2.size(); r++) {
            if (matches == 26) return true;

            int indexR = s2[r] - 'a';
            s2Count[indexR]++;
            if (s2Count[indexR] == s1Count[indexR]) matches++;
            else if (s2Count[indexR] == s1Count[indexR] + 1) matches--;

            int indexL = s2[l] - 'a';
            s2Count[indexL]--;
            if (s2Count[indexL] == s1Count[indexL]) matches++;
            else if (s2Count[indexL] == s1Count[indexL] - 1) matches--;

            l++;
        }

        return matches == 26;
    }
};