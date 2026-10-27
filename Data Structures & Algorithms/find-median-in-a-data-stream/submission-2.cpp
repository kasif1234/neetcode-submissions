#include <queue>
#include <vector>
using namespace std;

class MedianFinder {
    priority_queue<int> small;                             // max-heap
    priority_queue<int, vector<int>, greater<int>> large;  // min-heap

public:
    MedianFinder() {}

    void addNum(int num) {
        small.push(num);
        large.push(small.top());
        small.pop();

        if (large.size() > small.size()) {
            small.push(large.top());
            large.pop();
        }
    }

    double findMedian() {
        if (small.size() > large.size()) {
            return small.top();
        }
        return (small.top() + large.top()) / 2.0;
    }
};