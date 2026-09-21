

type item struct {
	d int
	p []int
}
type maxHeap []item

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].d > h[j].d } // max-heap
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)         { *h = append(*h, x.(item)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &maxHeap{}
	for _, p := range points {
		heap.Push(h, item{p[0]*p[0] + p[1]*p[1], p})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([][]int, 0, k)
	for _, it := range *h {
		res = append(res, it.p)
	}
	return res
}