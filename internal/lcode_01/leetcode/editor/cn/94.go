//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1360 👎 0

package main

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	cur := root
	var ans []int
	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		if stack.Len() != 0 {
			elem := stack.Back()
			top := elem.Value.(*TreeNode)
			ans = append(ans, top.Val)
			cur = top.Right
			stack.Remove(elem)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
