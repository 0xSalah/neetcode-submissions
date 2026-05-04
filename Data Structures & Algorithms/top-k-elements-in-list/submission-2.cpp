class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {

        map<int,int> count;
        int max = 0;

        for(int i =0; i< nums.size(); i++){
            count[nums[i]]++;
            if (max < count[nums[i]])
                max = count[nums[i]];
        }

        vector<int> arr;
        for(int i =max; i >= 0; i--){
            for(auto const& [num, freq] : count){
                if(freq == i)
                    arr.push_back(num);
                if (arr.size() == k)
                    return arr;
            }
        }
    }
};
