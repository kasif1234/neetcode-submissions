class Solution {
    public int leastInterval(char[] tasks, int n) {
        int[] count = new int[26];
        for (char t : tasks) count[t - 'A']++;

        int maxFreq = 0;
        for (int c : count) maxFreq = Math.max(maxFreq, c);

        int maxCount = 0;
        for (int c : count) if (c == maxFreq) maxCount++;

        return Math.max(tasks.length, (maxFreq - 1) * (n + 1) + maxCount);
    }
}