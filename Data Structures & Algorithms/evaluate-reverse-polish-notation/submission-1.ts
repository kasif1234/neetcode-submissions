class Solution {
    evalRPN(tokens: string[]): number {
        const stack: number[] = [];
        const ops = new Set(["+", "-", "*", "/"]);

        for (const tok of tokens) {
            if (ops.has(tok)) {
                const b = stack.pop()!;
                const a = stack.pop()!;
                let res: number;
                switch (tok) {
                    case "+": res = a + b; break;
                    case "-": res = a - b; break;
                    case "*": res = a * b; break;
                    default:  res = Math.trunc(a / b);
                }
                stack.push(res);
            } else {
                stack.push(parseInt(tok, 10));
            }
        }

        return stack[0];
    }
}