func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    indices := make([]int, n)
    for i := range indices {
        indices[i] = i
    }
    
    sort.Slice(indices, func(a, b int) bool {
        return position[indices[a]] > position[indices[b]]
    })
    
    stack := []float64{}
    
    for _, i := range indices {
        time := float64(target-position[i]) / float64(speed[i])
        if len(stack) == 0 || time > stack[len(stack)-1] {
            stack = append(stack, time)
        }
    }
    
    return len(stack)
}