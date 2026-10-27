import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MedianFinder struct {
    small *IntHeap // max-heap (negatives)
    large *IntHeap // min-heap
}

func Constructor() MedianFinder {
    return MedianFinder{small: &IntHeap{}, large: &IntHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
    heap.Push(this.small, -num)
    heap.Push(this.large, -heap.Pop(this.small).(int))

    if this.large.Len() > this.small.Len() {
        heap.Push(this.small, -heap.Pop(this.large).(int))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len() > this.large.Len() {
        return float64(-(*this.small)[0])
    }
    return float64(-(*this.small)[0]+(*this.large)[0]) / 2.0
}