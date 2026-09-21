class Solution {
    kClosest(points: number[][], k: number): number[][] {
        return points
            .slice()
            .sort((a, b) => (a[0]*a[0] + a[1]*a[1]) - (b[0]*b[0] + b[1]*b[1]))
            .slice(0, k);
    }
}