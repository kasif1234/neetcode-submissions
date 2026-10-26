import "container/heap"

type item struct{ time, tweetId, userId, idx int }

type maxHeap []item

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].time > h[j].time } // newest first
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *maxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type Twitter struct {
    time      int
    tweets    map[int][][2]int      // user -> [time, tweetId]
    following map[int]map[int]bool
}

func Constructor() Twitter {
    return Twitter{
        tweets:    map[int][][2]int{},
        following: map[int]map[int]bool{},
    }
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
    this.tweets[userId] = append(this.tweets[userId], [2]int{this.time, tweetId})
    this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    h := &maxHeap{}
    add := func(uid int) {
        t := this.tweets[uid]
        if len(t) > 0 {
            i := len(t) - 1
            *h = append(*h, item{t[i][0], t[i][1], uid, i - 1})
        }
    }

    add(userId)
    for f := range this.following[userId] {
        add(f)
    }
    heap.Init(h)

    res := []int{}
    for h.Len() > 0 && len(res) < 10 {
        top := heap.Pop(h).(item)
        res = append(res, top.tweetId)
        if top.idx >= 0 {
            t := this.tweets[top.userId][top.idx]
            heap.Push(h, item{t[0], t[1], top.userId, top.idx - 1})
        }
    }
    return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
    if this.following[followerId] == nil {
        this.following[followerId] = map[int]bool{}
    }
    this.following[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
    delete(this.following[followerId], followeeId)
}