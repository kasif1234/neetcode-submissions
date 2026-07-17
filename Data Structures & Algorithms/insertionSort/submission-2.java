/**
 * Definition for a pair.
 * class Pair {
 *     int key;
 *     String value;
 *     Pair() {}
 *     Pair(int key, String value) {
 *         this.key = key;
 *         this.value = value;
 *     }
 * }
 */
class Solution {
    public List<List<Pair>> insertionSort(List<Pair> pairs) {
        List<Pair> arr = new ArrayList<>(pairs);
        List<List<Pair>> states = new ArrayList<>();
        int n = arr.size();

        for (int i = 0; i < n; i++) {
            Pair current = arr.get(i);
            int j = i - 1;
            while (j >= 0 && arr.get(j).key > current.key) {
                arr.set(j + 1, arr.get(j));
                j--;
            }
            arr.set(j + 1, current);
            states.add(new ArrayList<>(arr));
        }

        return states;
    }
}