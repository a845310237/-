package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

var ans int
func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	depth(root)
	return ans -1
}

func depth(rt *TreeNode) int {
	temp := 0
	if rt == nil {
		return 0
	}
	L := depth(rt.Left)
	R := depth(rt.Right)
	if L + R + 1 > ans {
		ans = L + R + 1
	}
	if L > R {
		temp = L + 1
	} else {
		temp = R + 1
	}
	return temp
}

func main() {
	diameterOfBinaryTree()
}
