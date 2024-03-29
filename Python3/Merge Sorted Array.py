class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i = m - 1
        j = n - 1
        k = len(nums1) - 1

        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                k = k - 1
                i = i - 1
            else:
                nums1[k] = nums2[j]
                k = k - 1
                j = j - 1

'透過k從尾巴開始更新array的值
'透過雙指針去比對nums1 & nums2並同時控制k的位置