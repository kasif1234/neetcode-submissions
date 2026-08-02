class Solution {
    encode(strs: string[]): string {
        let res = "";
        for (const s of strs) {
            res += s.length + "#" + s;
        }
        return res;
    }

    decode(str: string): string[] {
        const res: string[] = [];
        let i = 0;
        while (i < str.length) {
            let j = i;
            while (str[j] !== "#") j++;
            const len = parseInt(str.substring(i, j));
            i = j + 1;
            res.push(str.substring(i, i + len));
            i += len;
        }
        return res;
    }
}