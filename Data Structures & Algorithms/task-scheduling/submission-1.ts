class Solution {
    leastInterval(tasks: string[], n: number): number {
        const count: number[] = new Array(26).fill(0);
        for (const t of tasks) count[t.charCodeAt(0) - 65]++;

        const maxFreq = Math.max(...count);
        const maxCount = count.filter(c => c === maxFreq).length;

        return Math.max(tasks.length, (maxFreq - 1) * (n + 1) + maxCount);
    }
}