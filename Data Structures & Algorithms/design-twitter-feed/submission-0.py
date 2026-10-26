import heapq
from collections import defaultdict

class Twitter:

    def __init__(self):
        self.time = 0
        self.tweets = defaultdict(list)    # userId -> [(time, tweetId), ...]
        self.following = defaultdict(set)  # userId -> {followeeIds}

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.tweets[userId].append((self.time, tweetId))
        self.time -= 1  # decreasing, so Python's min-heap acts as a max-heap

    def getNewsFeed(self, userId: int) -> List[int]:
        res = []
        heap = []

        for uid in self.following[userId] | {userId}:
            if self.tweets[uid]:
                idx = len(self.tweets[uid]) - 1
                t, tid = self.tweets[uid][idx]
                heap.append((t, tid, uid, idx - 1))

        heapq.heapify(heap)

        while heap and len(res) < 10:
            t, tid, uid, idx = heapq.heappop(heap)
            res.append(tid)
            if idx >= 0:
                t2, tid2 = self.tweets[uid][idx]
                heapq.heappush(heap, (t2, tid2, uid, idx - 1))

        return res

    def follow(self, followerId: int, followeeId: int) -> None:
        self.following[followerId].add(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.following[followerId].discard(followeeId)