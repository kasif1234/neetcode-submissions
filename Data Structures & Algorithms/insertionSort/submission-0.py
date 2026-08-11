def insertion_sort_states(pairs):
    arr = list(pairs)
    states = []
    n = len(arr)
    for i in range(n):
        key, val = arr[i]
        j = i - 1
        while j >= 0 and arr[j][0] > key:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = (key, val)
        states.append(list(arr))
    return states


# Example
pairs = [(5, "apple"), (2, "banana"), (9, "cherry")]
for state in insertion_sort_states(pairs):
    print(state)