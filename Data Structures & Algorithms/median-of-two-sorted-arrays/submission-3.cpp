class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        vector<int>& A = nums1.size() <= nums2.size() ? nums1 : nums2;
        vector<int>& B = nums1.size() <= nums2.size() ? nums2 : nums1;

        int m = A.size(), n = B.size();
        int total = m + n;
        int half = (total + 1) / 2;

        int lo = 0, hi = m;
        while (lo <= hi) {
            int i = (lo + hi) / 2;
            int j = half - i;

            int Aleft = (i > 0) ? A[i - 1] : INT_MIN;
            int Aright = (i < m) ? A[i] : INT_MAX;
            int Bleft = (j > 0) ? B[j - 1] : INT_MIN;
            int Bright = (j < n) ? B[j] : INT_MAX;

            if (Aleft <= Bright && Bleft <= Aright) {
                if (total % 2 == 1) {
                    return max(Aleft, Bleft);
                }
                return (max(Aleft, Bleft) + min(Aright, Bright)) / 2.0;
            } else if (Aleft > Bright) {
                hi = i - 1;
            } else {
                lo = i + 1;
            }
        }
        return 0.0;
    }
};