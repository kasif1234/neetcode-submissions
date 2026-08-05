class Solution {
    public int minEatingSpeed(int[] piles, int h) {
        int left = 1, right = 0;
        for (int pile : piles) {
            right = Math.max(right, pile);
        }
        int res = right;

        while (left <= right) {
            int k = left + (right - left) / 2;
            long hours = 0;
            for (int pile : piles) {
                hours += (pile + k - 1) / k; // ceil division
            }

            if (hours <= h) {
                res = k;
                right = k - 1;
            } else {
                left = k + 1;
            }
        }

        return res;
    }
}