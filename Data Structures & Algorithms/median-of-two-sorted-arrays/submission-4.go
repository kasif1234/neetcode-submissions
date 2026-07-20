func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2
    if len(A) > len(B) {
        A, B = B, A
    }

    m, n := len(A), len(B)
    total := m + n
    half := (total + 1) / 2

    lo, hi := 0, m
    for lo <= hi {
        i := (lo + hi) / 2
        j := half - i

        Aleft, Aright := math.MinInt32, math.MaxInt32
        Bleft, Bright := math.MinInt32, math.MaxInt32

        if i > 0 {
            Aleft = A[i-1]
        }
        if i < m {
            Aright = A[i]
        }
        if j > 0 {
            Bleft = B[j-1]
        }
        if j < n {
            Bright = B[j]
        }

        if Aleft <= Bright && Bleft <= Aright {
            if total%2 == 1 {
                return float64(max(Aleft, Bleft))
            }
            return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2.0
        } else if Aleft > Bright {
            hi = i - 1
        } else {
            lo = i + 1
        }
    }
    return 0.0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}