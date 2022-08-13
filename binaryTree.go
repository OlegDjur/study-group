package main

import "fmt"

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

// func (t *TreeNode) find(val int) *TreeNode {
// 	if t == nil {
// 		return nil
// 	}
// 	if val == t.val {

// 		tmp := t.delete(t.val)
// 		fmt.Println(t)
// 		return tmp
// 	}
// 	if val < t.val {
// 		t.left.find(val)
// 	} else if val > t.val {
// 		t.right.find(val)
// 	}
// 	return nil
// }

func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (t *TreeNode) remove(val int) *TreeNode {
	// var temp *TreeNode = nil

	if t.val == val {
		if t.left == nil && t.right == nil {

			t = nil
			fmt.Println(t)
			return t
		} else if t.left != nil && t.right == nil {
			tmp := t.left
			t = nil
			t = tmp

		} else if t.left == nil && t.right != nil {
			tmp := t.right
			t = nil
			t = tmp

		} else {
			tmp := t.right.min()
			t = nil
			t = tmp
		}
	} else if val < t.val {
		t.left.remove(val)
	} else {
		t.right.remove(val)
	}
	return nil
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
	tree.Delete(3)
	fmt.Println(tree.remove(3))
	tree.print()
}
