class Solution {
public:
    int carFleet(int target, vector<int>& position, vector<int>& speed) {
        int n = position.size();
        vector<int> indices(n);
        for (int i = 0; i < n; i++) indices[i] = i;
        
        sort(indices.begin(), indices.end(), [&](int a, int b) {
            return position[a] > position[b];
        });
        
        vector<double> stack;
        
        for (int i : indices) {
            double time = (double)(target - position[i]) / speed[i];
            if (stack.empty() || time > stack.back()) {
                stack.push_back(time);
            }
        }
        
        return stack.size();
    }
};