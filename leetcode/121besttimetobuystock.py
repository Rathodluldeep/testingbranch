from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ans=0
        minValueSoFar=prices[0]
        for i in range(1,len(prices)):
            profit=prices[i] - minValueSoFar
            if (profit > ans):
                ans=profit
            if (prices[i] < minValueSoFar):
                minValueSoFar=prices[i]
        return ans

prices=[7,1,5,3,6,4]
print(Solution().maxProfit(prices))