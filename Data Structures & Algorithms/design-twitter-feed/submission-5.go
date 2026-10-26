type feedItem struct{ time, tweetId, userId, idx int }

// move a node up until its parent is newer
func siftUp(h []feedItem, i int) {
    for i > 0 {
        p := (i - 1) / 2
        if h[p].time >= h[i].time {
            break
        }
        h[p], h[i] = h[i], h[p]
        i = p
    }
}

// move a node down until both children are older
func siftDown(h []feedItem, i int) {
    n := len(h)
    for {
        largest := i
        l, r := 2*i+1, 2*i+2
        if l < n && h[l].time > h[largest].time {
            largest = l
        }
        if r < n && h[r].time > h[largest].time {
            largest = r
        }
        if largest == i {
            break
        }
        h[i], h[largest] = h[largest], h[i]
        i = largest
    }
}

type Twitter struct {
    time      int
    tweets    map[int][][2]int // user -> [time, tweetId]
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
    h := []feedItem{}
    add := func(uid int) {
        t := this.tweets[uid]
        if len(t) > 0 {
            i := len(t) - 1
            h = append(h, feedItem{t[i][0], t[i][1], uid, i - 1})
        }
    }

    add(userId)
    for f := range this.following[userId] {
        add(f)
    }

    // build the heap
    for i := len(h)/2 - 1; i >= 0; i-- {
        siftDown(h, i)
    }

    res := []int{}
    for len(h) > 0 && len(res) < 10 {
        top := h[0]
        res = append(res, top.tweetId)

        if top.idx >= 0 {
            // replace the top with that user's next older tweet
            t := this.tweets[top.userId][top.idx]
            h[0] = feedItem{t[0], t[1], top.userId, top.idx - 1}
        } else {
            // that user has no more tweets: remove the top
            last := len(h) - 1
            h[0] = h[last]
            h = h[:last]
        }
        if len(h) > 0 {
            siftDown(h, 0)
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