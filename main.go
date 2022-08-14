package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) insert(val int) error {
	if t == nil {
		return fmt.Errorf("")
	} else if t.val > val {
		if t.left == nil {
			t.left = &TreeNode{val: val}
		} else {
			t.left.insert(val)
		}
	} else {
		if t.right == nil {
			t.right = &TreeNode{val: val}
		} else {
			t.right.insert(val)
		}
	}
	return nil
}

func (t *TreeNode) min() *TreeNode {
	if t.left != nil {
		t.left.min()
	} else {
		return t
	}
	return nil
}

func (t *TreeNode) max() {
	if t.right != nil {
		t.right.max()
	} else {
		fmt.Println(t.val)
		return
	}
}

func (t *TreeNode) print() {
	if t == nil {
		return
	}

	t.left.print()
	fmt.Println(t.val)
	t.right.print()
}

func (n *TreeNode) replaceNode(parent, replacement *TreeNode) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.left {
		parent.left = replacement
		return nil
	}
	parent.right = replacement
	return nil
}

func (t *TreeNode) remove(val int, parent *TreeNode) *TreeNode {
	if val < t.val {
		t.left.remove(val, t)
	} else if val > t.val {
		t.right.remove(val, t)
	} else if t.val == val {
		if t.left == nil && t.right == nil {
			t.replaceNode(parent, nil)
			return nil
		} else if t.left != nil && t.right == nil {
			t.replaceNode(parent, t.left)
			return nil
		} else if t.left == nil && t.right != nil {
			t.replaceNode(parent, t.right)
			return nil
		} else {
			tmp := t.right.min()
			l := t.left
			t.replaceNode(parent, tmp)
			tmp.left = l
			return nil
		}

	}
	return t
}

func main() {
	tree := &TreeNode{val: 8}

	tree.insert(4)
	tree.insert(2)
	tree.insert(3)
	tree.insert(10)
	tree.insert(6)
	tree.insert(7)

	tree.print()
	tree.remove(4, tree)
	fmt.Println()
	tree.print()
}
