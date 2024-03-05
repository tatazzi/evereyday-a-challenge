class Solution {
    public int missingNumber(int[] nums) {
        Arrays.sort(nums);
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != i) {
                return i;
            }
        }
        return nums.length;
    }
}
//Runtime: 6ms
//Memory: 45.03mb

class Solution {
    public int missingNumber(int[] nums) {
        int n = nums.length;
        int total = n*(n+1)/2;
        for(int i:nums)
            total -= i;
        return total;
    }
}
//Runtime: 0ms
//Memory: 44.96mb