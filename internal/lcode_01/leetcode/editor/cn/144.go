//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
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
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
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
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 784 👎 0
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := list.New()
	var ans []int
	cur := root

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			ans = append(ans, cur.Val)
			stack.PushBack(cur)
			cur = cur.Left
		}

		if stack.Len() != 0 {
			elem := stack.Back()
			cur = elem.Value.(*TreeNode).Right
			stack.Remove(elem)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	ans = append(ans, root.Val)

	tmp := preorderTraversal(root.Left)
	for _, t := range tmp {
		ans = append(ans, t)
	}

	tmp = preorderTraversal(root.Right)
	for _, t := range tmp {
		ans = append(ans, t)
	}

	return ans
}
