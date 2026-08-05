class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        int left = 1, right = *max_element(piles.begin(), piles.end());
        int res = right;

        while (left <= right) {
            int k = left + (right - left) / 2;
            long long hours = 0;
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
};