#1
class Solution:
    def countOdds(self, low: int, high: int) -> int:
        #the count of odd numbers between 1 and low - 1 is low / 2
        #the count of odd numbers between 1 and high is (high + 1 ) / 2
        return (high + 1) // 2 - low // 2

#2
class Solution:
  def countOdds(self, low: int, high: int) -> int:
    if low % 2 == 0:
      return (high-low+1)//2
    return (high-low)//2 + 1
    