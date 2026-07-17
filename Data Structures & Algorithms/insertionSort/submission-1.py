# Definition for a pair.
# class Pair:
#     def __init__(self, key: int, value: str):
#         self.key = key
#         self.value = value
class Solution:
    def insertionSort(self, pairs: List[Pair]) -> List[List[Pair]]:
        arr = list(pairs)
        states = []
        n = len(arr)

        for i in range(n):
            current = arr[i]
            j = i - 1
            while j >= 0 and arr[j].key > current.key:
                arr[j + 1] = arr[j]
                j -= 1
            arr[j + 1] = current
            states.append(list(arr))

        return states