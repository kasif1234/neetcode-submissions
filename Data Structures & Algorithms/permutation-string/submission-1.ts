class Solution {
    checkInclusion(s1: string, s2: string): boolean {
        if (s1.length > s2.length) return false;

        const s1Count = new Array(26).fill(0);
        const s2Count = new Array(26).fill(0);
        const a = 'a'.charCodeAt(0);

        for (let i = 0; i < s1.length; i++) {
            s1Count[s1.charCodeAt(i) - a]++;
            s2Count[s2.charCodeAt(i) - a]++;
        }

        let matches = 0;
        for (let i = 0; i < 26; i++) {
            if (s1Count[i] === s2Count[i]) matches++;
        }

        let l = 0;
        for (let r = s1.length; r < s2.length; r++) {
            if (matches === 26) return true;

            const indexR = s2.charCodeAt(r) - a;
            s2Count[indexR]++;
            if (s2Count[indexR] === s1Count[indexR]) matches++;
            else if (s2Count[indexR] === s1Count[indexR] + 1) matches--;

            const indexL = s2.charCodeAt(l) - a;
            s2Count[indexL]--;
            if (s2Count[indexL] === s1Count[indexL]) matches++;
            else if (s2Count[indexL] === s1Count[indexL] - 1) matches--;

            l++;
        }

        return matches === 26;
    }
}