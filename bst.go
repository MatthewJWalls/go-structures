package main

type TreeNode struct {
	Value int
	left *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (this *BinarySearchTree) Length() int {

	if this.root == nil {
		return 0
	} else {
		return length(this.root.left) + 1 + length(this.root.right)
	}

}

func length(n *TreeNode) int {

	if n == nil {
		return 0
	} else {
		return length(n.left) + 1 + length(n.right)
	}

}

func (this *BinarySearchTree) Contains(val int) bool {

	if this.root == nil {
		return false
	} else {

		n := this.root

		for {
			
			if n.Value == val {
				return true
			} else if val < n.Value && n.left != nil {
				n = n.left
			} else if val > n.Value && n.right != nil {
				n = n.right
			} else {
				return false
			}

		}


	}
	

}

func (this *BinarySearchTree) Insert(val int) {

	if this.root == nil {

		this.root = &TreeNode{val, nil, nil}

	} else {

		n := this.root

		for {

			if val < n.Value && n.left == nil {
				// create a left leaf
				n.left = &TreeNode{val, nil, nil}
			} else if val > n.Value && n.right == nil {
				// create a right leaf
				n.right = &TreeNode{val, nil, nil}
			} else if val == n.Value {
				// already there
				return
			} else if val < n.Value {
				// go down left
				n = n.left
			} else if val > n.Value {
				// go down right
				n = n.right
			} else {
				panic("Should not get to here")
			}

		}

	}

}

func (this *BinarySearchTree) Delete(val int){

	if this.root.Value == val && this.root.left == nil && this.root.right == nil {
		this.root = nil
	} else {
		this.delete(nil, this.root, val)
	}

}

func (this *BinarySearchTree) delete(p *TreeNode, n *TreeNode, val int) {

	for n != nil {

		if val == n.Value {

			// this is the node to delete

			if n.left == nil && n.right == nil {
				// case 1. deleting leaf
				if n.Value < p.Value {
					p.left = nil
				} else {
					p.right = nil
				}
				return
			} else if n.left != nil && n.right != nil {
				// case 3. deleting with 2 children
				n.Value = n.right.Value
				this.delete(n, n.right, n.right.Value)
				return
			} else if n.left != nil {
				// case 2. deleting with only left child
				n.Value = n.left.Value
				n.right = n.left.right
				n.left  = n.left.left
				return
			} else if n.right != nil {
				// case 2. deleting with only right child
				n.Value = n.right.Value
				n.left  = n.right.left
				n.right = n.right.right
				return
			}
		} else if val < n.Value {
			// go left
			p = n
			n = n.left
		} else if val > n.Value {
			// go right
			p = n
			n = n.right
		} else {
			panic("Don't know what to do")
		}

	}

}

func NewBinarySearchTree() (*BinarySearchTree) {
	return &BinarySearchTree{}
}
