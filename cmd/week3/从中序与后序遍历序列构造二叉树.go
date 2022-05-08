package main

/**
106. 从中序与后序遍历序列构造二叉树

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
**/

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var fn func(int, int, int, int) *TreeNode

	fn = func(p1, end1, p2, end2 int) *TreeNode {
		//临界点为集合左右两边界错位了,那么说明集合没有元素了
		if p1 > end1 || p2 > end2 {
			return nil
		}

		v := postorder[end1]

		node := &TreeNode{
			Val: v,
		}

		//标记根节点在中序遍历数组位置
		var index int

		for i := p2; i <= end2; i++ {
			if v == inorder[i] {
				index = i

				break
			}
		}

		//右子树集合
		node.Right = fn(end1+index-end2, end1-1, index+1, end2)
		//左子树集合
		node.Left = fn(p1, end1+index-end2-1, p2, index-1)

		return node
	}

	return fn(0, len(postorder)-1, 0, len(inorder)-1)
}
