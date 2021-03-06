package main

/**
* https://leetcode.com/problems/binary-tree-level-order-traversal/description/
*
* Example :
* Given binary tree [3,9,20,null,null,15,7],
*     3
*    / \
*   9  20
*     /  \
*    15   7
* return its level order traversal as:
* [
*   [3],
*   [9,20],
*   [15,7]
* ]
*
* 34 / 34 test cases passed.
* Runtime: 4 ms
* Beats 100%
* https://leetcode.com/submissions/detail/173475146/
**/

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret, stack := [][]int{}, []*TreeNode{root}
	for len(stack) > 0 {
		lenS := len(stack)
		var list []int
		for i := 0; i < lenS; i++ {
			n := stack[i]
			list = append(list, n.Val)
			if n.Left != nil {
				stack = append(stack, n.Left)
			}
			if n.Right != nil {
				stack = append(stack, n.Right)
			}
		}
		stack = stack[lenS:]
		ret = append(ret, list)
	}
	return ret
}
