class Solution:
    # 594. 最长和谐子序列
    def findLHS(self, nums: List[int]) -> int:
        nums.sort()
        n = len(nums)
        if n == 1 or nums[0] == nums[-1]:
            return 0
        l, r = 0, 1
        ans = 0
        while l < n:
            while r < n and nums[r] - nums[l] <= 1:
                r += 1
            if nums[l] != nums[r - 1]:
                ans = max(ans, r - l)
            l += 1
        return ans