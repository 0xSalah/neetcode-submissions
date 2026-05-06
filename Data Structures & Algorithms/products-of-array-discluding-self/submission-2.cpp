class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> result;

        for(int i = 0; i < nums.size(); i++){
            int x = i;
            int mult = 1;
            for (int j = 0; j < nums.size(); j++){
                if (j != i){
                    mult *= nums[j];
                }
            }
            result.push_back(mult);
        }
        return result;
    }
};
