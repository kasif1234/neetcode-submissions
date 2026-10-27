class MinHeap {
    private data: number[] = [];

    size(): number { return this.data.length; }
    peek(): number { return this.data[0]; }

    push(val: number): void {
        this.data.push(val);
        let i = this.data.length - 1;
        while (i > 0) {
            const p = (i - 1) >> 1;
            if (this.data[p] <= this.data[i]) break;
            [this.data[p], this.data[i]] = [this.data[i], this.data[p]];
            i = p;
        }
    }

    pop(): number {
        const top = this.data[0];
        const last = this.data.pop()!;
        if (this.data.length > 0) {
            this.data[0] = last;
            let i = 0;
            const n = this.data.length;
            while (true) {
                let smallest = i;
                const l = 2 * i + 1, r = 2 * i + 2;
                if (l < n && this.data[l] < this.data[smallest]) smallest = l;
                if (r < n && this.data[r] < this.data[smallest]) smallest = r;
                if (smallest === i) break;
                [this.data[smallest], this.data[i]] = [this.data[i], this.data[smallest]];
                i = smallest;
            }
        }
        return top;
    }
}

class MedianFinder {
    private small = new MinHeap(); // max-heap (negatives)
    private large = new MinHeap(); // min-heap

    constructor() {}

    addNum(num: number): void {
        this.small.push(-num);
        this.large.push(-this.small.pop());

        if (this.large.size() > this.small.size()) {
            this.small.push(-this.large.pop());
        }
    }

    findMedian(): number {
        if (this.small.size() > this.large.size()) {
            return -this.small.peek();
        }
        return (-this.small.peek() + this.large.peek()) / 2;
    }
}