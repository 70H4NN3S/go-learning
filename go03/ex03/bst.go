package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func (b *BST) Insert(val int) {
	if b.Root == nil {
		b.Root = &TreeNode{val, nil, nil}
		return
	}

	for n := b.Root; ; {
		if n.Val == val {
			return
		} else if n.Val > val {
			if n.Left == nil {
				n.Left = &TreeNode{val, nil, nil}
				return
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = &TreeNode{val, nil, nil}
				return
			}
			n = n.Right
		}
	}
}

func (b BST) Search(val int) bool {
	for n := b.Root; n != nil; {
		if n.Val == val {
			return true
		} else if n.Val > val {
			n = n.Left
		} else if n.Val < val {
			n = n.Right
		}
	}
	return false
}

func (b BST) ToSlice() []int {
	result := make([]int, 0)
	stack := []TreeNode{*b.Root}

	for n := b.Root; len(stack) != 0; {
		if n.Left != nil {
			stack = append(stack, *n.Left)
			n.Left = nil
			n = &stack[len(stack)-1]
			continue
		} else if n.Right != nil {
			result = append(result, n.Val)
			stack = stack[:len(stack)-1]
			stack = append(stack, *n.Right)
			n.Right = nil
			n = &stack[len(stack)-1]
			continue
		} else {
			result = append(result, n.Val)
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			n = &stack[len(stack)-1]
		}
	}
	return result
}
