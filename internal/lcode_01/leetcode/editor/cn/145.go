////给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
////
////
////
//// 示例 1：
////
////
////输入：root = [1,null,2,3]
////输出：[3,2,1]
////
////
//// 示例 2：
////
////
////输入：root = []
////输出：[]
////
////
//// 示例 3：
////
////
////输入：root = [1]
////输出：[1]
////
////
////
////
//// 提示：
////
////
//// 树中节点的数目在范围 [0, 100] 内
//// -100 <= Node.val <= 100
////
////
////
////
//// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//// Related Topics 栈 树 深度优先搜索 二叉树 👍 801 👎 0
//
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		ans        []int
		preVisited *TreeNode
	)
	cur := root
	stack := list.New()

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		elem := stack.Back()
		top := elem.Value.(*TreeNode)

		// 没有子节点
		if (top.Left == nil && top.Right == nil) ||
			// 没有右节点且左节点已经访问过
			(top.Right == nil && preVisited == top.Left) ||
			// 右节点已经访问过
			preVisited == top.Right {
			ans = append(ans, top.Val)
			stack.Remove(elem)
			preVisited = top
		} else {
			cur = top.Right
		}

	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	tmp := postorderTraversal(root.Left)
	for _, t := range tmp {
		ans = append(ans, t)
	}

	tmp = postorderTraversal(root.Right)
	for _, t := range tmp {
		ans = append(ans, t)
	}

	return append(ans, root.Val)
}

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	var (
		ans        []int
		preVisited *TreeNode
	)
	cur := root

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		elem := stack.Back()
		top := elem.Value.(*TreeNode)

		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
			ans = append(ans, top.Val)
			stack.Remove(elem)
			preVisited = top
		} else {
			cur = top.Right
		}
	}

	return ans
}
