class Twitter {
    private int time = 0;
    private Map<Integer, List<int[]>> tweets = new HashMap<>();   // user -> [time, tweetId]
    private Map<Integer, Set<Integer>> following = new HashMap<>();

    public Twitter() {}

    public void postTweet(int userId, int tweetId) {
        tweets.computeIfAbsent(userId, k -> new ArrayList<>()).add(new int[]{time++, tweetId});
    }

    public List<Integer> getNewsFeed(int userId) {
        List<Integer> res = new ArrayList<>();
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[0] - a[0]); // newest first

        Set<Integer> users = new HashSet<>(following.getOrDefault(userId, new HashSet<>()));
        users.add(userId);

        for (int uid : users) {
            List<int[]> list = tweets.get(uid);
            if (list != null && !list.isEmpty()) {
                int i = list.size() - 1;
                pq.add(new int[]{list.get(i)[0], list.get(i)[1], uid, i - 1});
            }
        }

        while (!pq.isEmpty() && res.size() < 10) {
            int[] top = pq.poll();
            res.add(top[1]);
            if (top[3] >= 0) {
                int[] t = tweets.get(top[2]).get(top[3]);
                pq.add(new int[]{t[0], t[1], top[2], top[3] - 1});
            }
        }
        return res;
    }

    public void follow(int followerId, int followeeId) {
        following.computeIfAbsent(followerId, k -> new HashSet<>()).add(followeeId);
    }

    public void unfollow(int followerId, int followeeId) {
        if (following.containsKey(followerId)) following.get(followerId).remove(followeeId);
    }
}