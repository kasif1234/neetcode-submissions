class MedianFinder {
    private java.util.PriorityQueue<Integer> small; // max-heap
    private java.util.PriorityQueue<Integer> large; // min-heap

    public MedianFinder() {
        small = new java.util.PriorityQueue<>(java.util.Collections.reverseOrder());
        large = new java.util.PriorityQueue<>();
    }

    public void addNum(int num) {
        small.offer(num);
        large.offer(small.poll());

        if (large.size() > small.size()) {
            small.offer(large.poll());
        }
    }

    public double findMedian() {
        if (small.size() > large.size()) {
            return small.peek();
        }
        return (small.peek() + large.peek()) / 2.0;
    }
}