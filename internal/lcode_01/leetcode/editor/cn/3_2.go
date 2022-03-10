//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口 👍 7076 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
// start 不动 end 向后移动
// 当end 遇到重复字符 start 应该放在上一个重复字符位置的最后一位
// 不管字符重复不重复 都要更新end 字符的位置 和 长度
// 怎样判断是否遇到重复字符，且怎么知道上一个重复字符的位置？--用哈希字典的key来判断是否重复，用value来记录该字符的下一个不重复的位置。
func lengthOfLongestSubstring2(s string) int {

	strIndex := make(map[string]int) // key为字符 value 为字符串位置
	start, end, result := 0, 0, 0
	for end < len(s) {

		if index, ok := strIndex[string(s[end])]; ok {
			start = max2(start, index+1) // start 记录下一个不重复的字符的位置
		}

		strIndex[string(s[end])] = end
		result = max2(result, end-start+1)
		end++
	}
	return result
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
