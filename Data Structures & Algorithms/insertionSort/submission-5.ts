/**
 * Pair class to store key-value pairs
 */
// class Pair {
//     /**
//      * @param {number} key The key to be stored in the pair
//      * @param {string} value The value to be stored in the pair
//      */
//     constructor(key, value) {
//         this.key = key;
//         this.value = value;
//     }
// }
class Solution {
    /**
     * @param {Pair[]} pairs
     * @returns {Pair[][]}
     */
    insertionSort(pairs: Pair[]): Pair[][] {
        const arr: Pair[] = [...pairs];
        const states: Pair[][] = [];
        const n = arr.length;

        for (let i = 0; i < n; i++) {
            const current = arr[i];
            let j = i - 1;
            while (j >= 0 && arr[j].key > current.key) {
                arr[j + 1] = arr[j];
                j--;
            }
            arr[j + 1] = current;
            states.push([...arr]);
        }

        return states;
    }
}