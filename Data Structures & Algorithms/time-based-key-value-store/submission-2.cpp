class TimeMap {
public:
    unordered_map<string, vector<pair<int, string>>> keyStore;

    TimeMap() {
        
    }

    void set(string key, string value, int timestamp) {
        keyStore[key].push_back({timestamp, value});
    }

    string get(string key, int timestamp) {
        string res = "";
        if (keyStore.find(key) == keyStore.end()) return res;

        auto& values = keyStore[key];
        int lo = 0, hi = values.size() - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (values[mid].first <= timestamp) {
                res = values[mid].second;
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
        return res;
    }
};