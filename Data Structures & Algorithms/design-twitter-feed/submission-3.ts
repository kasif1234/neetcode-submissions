class Twitter {
    private time = 0;
    private tweets = new Map<number, [number, number][]>();   // user -> [time, tweetId]
    private following = new Map<number, Set<number>>();

    constructor() {}

    postTweet(userId: number, tweetId: number): void {
        if (!this.tweets.has(userId)) this.tweets.set(userId, []);
        this.tweets.get(userId)!.push([this.time++, tweetId]);
    }

    getNewsFeed(userId: number): number[] {
        const users = new Set(this.following.get(userId) ?? []);
        users.add(userId);

        // pointer to each user's newest unused tweet
        const ptr = new Map<number, number>();
        for (const u of users) ptr.set(u, (this.tweets.get(u)?.length ?? 0) - 1);

        const res: number[] = [];
        while (res.length < 10) {
            let best = -1, bestTime = -1;
            for (const [u, i] of ptr) {
                if (i >= 0) {
                    const t = this.tweets.get(u)![i][0];
                    if (t > bestTime) { bestTime = t; best = u; }
                }
            }
            if (best === -1) break;
            const i = ptr.get(best)!;
            res.push(this.tweets.get(best)![i][1]);
            ptr.set(best, i - 1);
        }
        return res;
    }

    follow(followerId: number, followeeId: number): void {
        if (!this.following.has(followerId)) this.following.set(followerId, new Set());
        this.following.get(followerId)!.add(followeeId);
    }

    unfollow(followerId: number, followeeId: number): void {
        this.following.get(followerId)?.delete(followeeId);
    }
}