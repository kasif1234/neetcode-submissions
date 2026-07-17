/**
 * Definition for a pair.
 * type Pair struct {
 *     Key   int
 *     Value string
 * }
 */
func insertionSort(pairs []Pair) [][]Pair {
    arr := make([]Pair, len(pairs))
    copy(arr, pairs)
    n := len(arr)
    states := make([][]Pair, 0, n)

    for i := 0; i < n; i++ {
        current := arr[i]
        j := i - 1
        for j >= 0 && arr[j].Key > current.Key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = current

        stateCopy := make([]Pair, n)
        copy(stateCopy, arr)
        states = append(states, stateCopy)
    }

    return states
}