class Twitter {
    int time = 0;
    unordered_map<int, vector<pair<int,int>>> tweets;   // user -> (time, tweetId)
    unordered_map<int, unordered_set<int>> following;

public:
    Twitter() {}

    void postTweet(int userId, int tweetId) {
        tweets[userId].push_back({time++, tweetId});
    }

    vector<int> getNewsFeed(int userId) {
        priority_queue<array<int,4>> pq;   // max-heap on time (first element)

        auto add = [&](int uid) {
            auto it = tweets.find(uid);
            if (it == tweets.end() || it->second.empty()) return;
            int i = it->second.size() - 1;
            pq.push({it->second[i].first, it->second[i].second, uid, i - 1});
        };

        add(userId);
        for (int f : following[userId]) add(f);

        vector<int> res;
        while (!pq.empty() && res.size() < 10) {
            auto [t, tid, uid, i] = pq.top();
            pq.pop();
            res.push_back(tid);
            if (i >= 0) {
                auto [t2, tid2] = tweets[uid][i];
                pq.push({t2, tid2, uid, i - 1});
            }
        }
        return res;
    }

    void follow(int followerId, int followeeId) {
        following[followerId].insert(followeeId);
    }

    void unfollow(int followerId, int followeeId) {
        following[followerId].erase(followeeId);
    }
};