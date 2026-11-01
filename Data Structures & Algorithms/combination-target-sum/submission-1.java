class Solution {
    public List<List<Integer>> combinationSum(int[] nums, int target) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(nums, target, 0, 0, new ArrayList<>(), res);
        return res;
    }

    private void dfs(int[] nums, int target, int start, int total,
                     List<Integer> cur, List<List<Integer>> res) {
        if (total == target) {
            res.add(new ArrayList<>(cur));   // save a copy
            return;
        }
        for (int j = start; j < nums.length; j++) {
            if (total + nums[j] > target) continue;
            cur.add(nums[j]);
            dfs(nums, target, j, total + nums[j], cur, res);  // j: reuse allowed
            cur.remove(cur.size() - 1);      // backtrack
        }
    }
}