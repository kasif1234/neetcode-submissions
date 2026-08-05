class Solution {
    minEatingSpeed(piles: number[], h: number): number {
        let left = 1;
        let right = Math.max(...piles);
        let res = right;

        while (left <= right) {
            const k = Math.floor((left + right) / 2);
            let hours = 0;
            for (const pile of piles) {
                hours += Math.ceil(pile / k);
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