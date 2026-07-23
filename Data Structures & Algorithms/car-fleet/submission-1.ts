class Solution {
    carFleet(target: number, position: number[], speed: number[]): number {
        const n = position.length;
        const indices = Array.from({ length: n }, (_, i) => i);
        indices.sort((a, b) => position[b] - position[a]);
        
        const stack: number[] = [];
        
        for (const i of indices) {
            const time = (target - position[i]) / speed[i];
            if (stack.length === 0 || time > stack[stack.length - 1]) {
                stack.push(time);
            }
        }
        
        return stack.length;
    }
}