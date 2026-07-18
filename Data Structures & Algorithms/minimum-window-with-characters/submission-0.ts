class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {string}
     */
    minWindow(s: string, t: string): string {
        if (t.length === 0 || s.length < t.length) return "";

        const need = new Map<string, number>();
        for (const c of t) {
            need.set(c, (need.get(c) || 0) + 1);
        }

        const window = new Map<string, number>();
        let have = 0;
        const required = need.size;

        let resLen = Infinity;
        let resLeft = 0;

        let left = 0;
        for (let right = 0; right < s.length; right++) {
            const c = s[right];
            window.set(c, (window.get(c) || 0) + 1);

            if (need.has(c) && window.get(c) === need.get(c)) {
                have++;
            }

            while (have === required) {
                if (right - left + 1 < resLen) {
                    resLen = right - left + 1;
                    resLeft = left;
                }

                const leftChar = s[left];
                window.set(leftChar, window.get(leftChar)! - 1);
                if (need.has(leftChar) && window.get(leftChar)! < need.get(leftChar)!) {
                    have--;
                }
                left++;
            }
        }

        return resLen === Infinity ? "" : s.substring(resLeft, resLeft + resLen);
    }
}