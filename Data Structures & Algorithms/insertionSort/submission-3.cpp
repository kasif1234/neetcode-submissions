/**
 * Definition for a pair.
 * struct Pair {
 *     int key;
 *     string value;
 *     Pair() {}
 *     Pair(int key, string value) : key(key), value(value) {}
 * };
 */
class Solution {
public:
    vector<vector<Pair>> insertionSort(vector<Pair>& pairs) {
        vector<Pair> arr = pairs;
        vector<vector<Pair>> states;
        int n = static_cast<int>(arr.size());

        for (int i = 0; i < n; i++) {
            Pair current = arr[i];
            int j = i - 1;
            while (j >= 0 && arr[j].key > current.key) {
                arr[j + 1] = arr[j];
                j--;
            }
            arr[j + 1] = current;
            states.push_back(arr);
        }

        return states;
    }
};