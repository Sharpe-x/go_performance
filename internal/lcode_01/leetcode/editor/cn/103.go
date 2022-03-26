//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 广度优先搜索 二叉树 👍 606 👎 0
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	search(root, 0, &res)
	return res
}
func search(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}

	search(root.Left, depth+1, res)
	search(root.Right, depth+1, res)
}

//leetcode submit region end(Prohibit modification and deletion)
