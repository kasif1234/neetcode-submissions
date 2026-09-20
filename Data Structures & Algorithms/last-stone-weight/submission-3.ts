class Solution {
    lastStoneWeight(stones: number[]): number {
        const heap: number[] = [];

        const swap = (a: number, b: number) => {
            [heap[a], heap[b]] = [heap[b], heap[a]];
        };

        const push = (v: number) => {
            heap.push(v);
            let i = heap.length - 1;
            while (i > 0) {
                const p = (i - 1) >> 1;
                if (heap[p] >= heap[i]) break;
                swap(p, i);
                i = p;
            }
        };

        const pop = (): number => {
            const top = heap[0];
            const last = heap.pop()!;
            if (heap.length > 0) {
                heap[0] = last;
                let i = 0;
                while (true) {
                    const l = 2 * i + 1, r = l + 1;
                    let m = i;
                    if (l < heap.length && heap[l] > heap[m]) m = l;
                    if (r < heap.length && heap[r] > heap[m]) m = r;
                    if (m === i) break;
                    swap(m, i);
                    i = m;
                }
            }
            return top;
        };

        for (const s of stones) push(s);

        while (heap.length > 1) {
            const y = pop();
            const x = pop();
            if (y > x) push(y - x);
        }

        return heap.length ? heap[0] : 0;
    }
}