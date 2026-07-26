class TimeMap:

    def __init__(self):
        self.keyStore = {}  # key -> list of [timestamp, value]

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.keyStore:
            self.keyStore[key] = []
        self.keyStore[key].append([timestamp, value])

    def get(self, key: str, timestamp: int) -> str:
        values = self.keyStore.get(key, [])
        res = ""
        lo, hi = 0, len(values) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if values[mid][0] <= timestamp:
                res = values[mid][1]
                lo = mid + 1
            else:
                hi = mid - 1
        return res