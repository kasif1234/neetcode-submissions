func evalRPN(tokens []string) int {
    stack := make([]int, 0, len(tokens))

    push := func(x int) { stack = append(stack, x) }
    pop := func() int {
        n := len(stack) - 1
        v := stack[n]
        stack = stack[:n]
        return v
    }

    for _, tok := range tokens {
        switch tok {
        case "+", "-", "*", "/":
            b := pop()
            a := pop()
            switch tok {
            case "+":
                push(a + b)
            case "-":
                push(a - b)
            case "*":
                push(a * b)
            case "/":
                push(a / b) // Go integer division truncates toward zero
            }
        default:
            n, _ := strconv.Atoi(tok)
            push(n)
        }
    }

    return pop()
}