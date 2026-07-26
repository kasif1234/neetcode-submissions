type Pair struct {
    timestamp int
    value     string
}

type TimeMap struct {
    keyStore map[string][]Pair
}

func Constructor() TimeMap {
    return TimeMap{keyStore: make(map[string][]Pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.keyStore[key] = append(this.keyStore[key], Pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    values, ok := this.keyStore[key]
    res := ""
    if !ok {
        return res
    }

    lo, hi := 0, len(values)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if values[mid].timestamp <= timestamp {
            res = values[mid].value
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return res
}