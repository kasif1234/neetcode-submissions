class TimeMap {

    private Map<String, List<Pair>> keyStore;

    private static class Pair {
        int timestamp;
        String value;
        Pair(int timestamp, String value) {
            this.timestamp = timestamp;
            this.value = value;
        }
    }

    public TimeMap() {
        keyStore = new HashMap<>();
    }

    public void set(String key, String value, int timestamp) {
        keyStore.computeIfAbsent(key, k -> new ArrayList<>()).add(new Pair(timestamp, value));
    }

    public String get(String key, int timestamp) {
        List<Pair> values = keyStore.getOrDefault(key, new ArrayList<>());
        String res = "";
        int lo = 0, hi = values.size() - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (values.get(mid).timestamp <= timestamp) {
                res = values.get(mid).value;
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
        return res;
    }
}