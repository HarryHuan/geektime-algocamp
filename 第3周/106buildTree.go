package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var in []int
var post []int

func buildTree(inorder []int, postorder []int) *TreeNode {
	in = inorder
	post = postorder
	return build(0, len(in)-1, 0, len(post)-1)
}

func build(l1, r1, l2, r2 int) *TreeNode {
	if l2 > r2 {
		return nil
	}
	root := &TreeNode{Val: post[r2]}
	mid := l1
	for in[mid] != root.Val {
		mid++
	}
	root.Left = build(l1, mid-1, l2, mid-1-l1+l2)
	root.Right = build(mid+1, r1, mid-l1+l2, r2-1)
	return root
}
