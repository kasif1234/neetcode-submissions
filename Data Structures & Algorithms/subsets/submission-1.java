class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(0, nums, new ArrayList<>(), res);
        return res;
    }

    private void dfs(int i, int[] nums, List<Integer> subset, List<List<Integer>> res) {
        if (i == nums.length) {
            res.add(new ArrayList<>(subset)); // save a copy
            return;
        }

        // choice 1: take nums[i]
        subset.add(nums[i]);
        dfs(i + 1, nums, subset, res);

        // undo, then choice 2: skip nums[i]
        subset.remove(subset.size() - 1);
        dfs(i + 1, nums, subset, res);
    }
}