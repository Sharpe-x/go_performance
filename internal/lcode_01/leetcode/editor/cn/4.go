//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
// Related Topics 数组 二分查找 分治 👍 5251 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 && n == 0 {
		return float64(0)
	}

	if m == 0 {
		if n%2 == 0 {
			return float64(nums2[n/2-1]+nums2[n/2]) / 2.0
		} else {
			return float64(nums2[n/2])
		}
	}

	if n == 0 {
		if m%2 == 0 {
			return float64(nums1[m/2-1]+nums1[m/2]) / 2.0
		} else {
			return float64(nums1[m/2])
		}
	}

	nums3 := make([]int, m+n)
	i, j, count := 0, 0, 0
	for count != m+n {
		if i == m {
			for j != n {
				nums3[count] = nums2[j]
				count++
				j++
			}
			break
		}

		if j == n {
			for i != m {
				nums3[count] = nums1[i]
				count++
				i++
			}
			break
		}

		if nums1[i] < nums2[j] {
			nums3[count] = nums1[i]
			i++
			count++
		} else {
			nums3[count] = nums2[j]
			j++
			count++
		}
	}

	if count%2 == 0 {
		return float64(nums3[count/2-1]+nums3[count/2]) / 2.0
	}

	return float64(nums3[count/2])
}

//leetcode submit region end(Prohibit modification and deletion)
