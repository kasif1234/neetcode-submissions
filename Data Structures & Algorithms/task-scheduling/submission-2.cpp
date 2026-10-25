class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        vector<int> count(26, 0);
        for (char t : tasks) count[t - 'A']++;

        int maxFreq = *max_element(count.begin(), count.end());
        int maxCount = 0;
        for (int c : count) if (c == maxFreq) maxCount++;

        return max((int)tasks.size(), (maxFreq - 1) * (n + 1) + maxCount);
    }
};