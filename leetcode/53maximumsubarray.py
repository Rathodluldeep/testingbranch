'''give the interger array nums and find the subarraay with largest sun dn return its sum'''
from typing import List


class Solution:
    def maxSubArray(self, nums:List[int]) -> int:
        maxSoFar =nums[0]
        currentSum=nums[0]
        for i in range(1,len(nums)):
            if (currentSum<0):
                currentSum=0
            currentSum +=nums[i]
            if(currentSum> maxSoFar):
                maxSoFar=currentSum

        return maxSoFar