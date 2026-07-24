class Solution {
    dailyTemperatures(temperatures: number[]): number[] {
        const n = temperatures.length;
        const result: number[] = new Array(n).fill(0);
        const stack: number[] = []; // stores indices

        for (let i = 0; i < n; i++) {
            while (stack.length > 0 && temperatures[i] > temperatures[stack[stack.length - 1]]) {
                const prevIndex = stack.pop()!;
                result[prevIndex] = i - prevIndex;
            }
            stack.push(i);
        }

        return result;
    }
}