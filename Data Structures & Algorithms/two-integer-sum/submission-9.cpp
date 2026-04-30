class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> sol;

        for(int i = 0; i < nums.size(); i++){
            for (int y = (nums.size()-1) ; y >0; y--){
                if (nums[i] + nums[y] == target && i != y){
                    sol.push_back(i);
                    sol.push_back(y);
                    return sol;
                }
            }
        }
    return {};  
    }
};
