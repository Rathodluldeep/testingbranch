
from typing import List


class Solutions:
    def findMaxConsecutiveOnes(self, nums:List[int]) -> int:
        ans=0
        currenCount=0
        for i in range(len(nums)):
            if(nums[i]==0):
                currenCount=0
            else:
                currenCount += 1
            
            if(currenCount>ans):
                ans=currenCount
        return currenCount
    

nums=[1,1,0,1,1,1,0,1,1,0,1,0,1,0,1,1,1,1]
print(Solutions().findMaxConsecutiveOnes(nums))