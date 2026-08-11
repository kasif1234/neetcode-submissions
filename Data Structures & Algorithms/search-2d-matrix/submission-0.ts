class Solution {
    /**
     * @param {number[][]} matrix
     * @param {number} target
     * @return {boolean}
     */
    searchMatrix(matrix: number[][], target: number): boolean {
        const m = matrix.length, n = matrix[0].length;
        let lo = 0, hi = m * n - 1;

        while (lo <= hi) {
            const mid = Math.floor((lo + hi) / 2);
            const row = Math.floor(mid / n), col = mid % n;
            const val = matrix[row][col];

            if (val === target) {
                return true;
            } else if (val < target) {
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }

        return false;
    }
}