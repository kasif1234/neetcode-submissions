class Solution {
    public String minWindow(String s, String t) {
        if (t.length() == 0 || s.length() < t.length()) return "";

        Map<Character, Integer> need = new HashMap<>();
        for (char c : t.toCharArray()) {
            need.put(c, need.getOrDefault(c, 0) + 1);
        }

        Map<Character, Integer> window = new HashMap<>();
        int have = 0;
        int required = need.size();

        int resLen = Integer.MAX_VALUE;
        int resLeft = 0;

        int left = 0;
        for (int right = 0; right < s.length(); right++) {
            char c = s.charAt(right);
            window.put(c, window.getOrDefault(c, 0) + 1);

            if (need.containsKey(c) && window.get(c).intValue() == need.get(c).intValue()) {
                have++;
            }

            while (have == required) {
                if (right - left + 1 < resLen) {
                    resLen = right - left + 1;
                    resLeft = left;
                }

                char leftChar = s.charAt(left);
                window.put(leftChar, window.get(leftChar) - 1);
                if (need.containsKey(leftChar) && window.get(leftChar) < need.get(leftChar)) {
                    have--;
                }
                left++;
            }
        }

        return resLen == Integer.MAX_VALUE ? "" : s.substring(resLeft, resLeft + resLen);
    }
}