package main

import "testing"

func TestInsertAndContains(t *testing.T) {

	tree := NewBinarySearchTree()

	tree.Insert(10)
	tree.Insert(8)
	tree.Insert(11)
	tree.Insert(2)
	tree.Insert(9)
	tree.Insert(12)

	if ! tree.Contains(10) {
		t.Error("Tree did not contain 3.")
	}

	if ! tree.Contains(8) {
		t.Error("Tree did not contain 1.")
	}

	if ! tree.Contains(11) {
		t.Error("Tree did not contain 4.")
	}

	if ! tree.Contains(2) {
		t.Error("Tree did not contain 5.")
	}

	if ! tree.Contains(9) {
		t.Error("Tree did not contain 9.")
	}

	if ! tree.Contains(12) {
		t.Error("Tree contained 2")
	}

	if tree.Length() != 6 {
		t.Error("Tree length() was incorrect")
	}

}

func TestTreeDelete(t *testing.T) {

	tree := NewBinarySearchTree()

	tree.Insert(10)
	tree.Insert(8)
	tree.Insert(11)
	tree.Insert(2)
	tree.Insert(9)
	tree.Insert(12)

	tree.Delete(8)

	if tree.Contains(8) {
		t.Error("Still contained deleted value.")
	}

	if tree.Length() != 5 {
		t.Error("Length was wrong")
	}

	tree.Delete(2)

	if tree.Length() != 4 {
		t.Error("Length was wrong")
	}

	tree.Delete(12)

	if tree.Length() != 3 {
		t.Error("Length was wrong")
	}

	if tree.root.Value != 10 {
		t.Error("Tree was incorrect.")
	}

	if tree.root.left.Value != 9 {
		t.Error("Tree was incorrect.")
	}

	if tree.root.right.Value != 11 {
		t.Error("Tree was incorrect.")
	}

	if tree.root.left.left != nil {
		t.Error("Tree was incorrect.")
	}

	if tree.root.left.right != nil {
		t.Error("Tree was incorrect.")
	}

	if tree.root.right.right != nil {
		t.Error("Tree was incorrect.")
	}

	if tree.root.right.left != nil {
		t.Error("Tree was incorrect.")
	}

	tree.Delete(9)
	tree.Delete(11)

	if tree.Length() != 1 {
		t.Error("Tree was incorrect.")
	}

	tree.Delete(10)

	if tree.Length() != 0 {
		t.Error("Tree was incorrect.")
	}

}

func TestBreadth(t *testing.T) {

	tree := NewBinarySearchTree()
	
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)

	nodes := tree.Breadth()

	if nodes.Length() != 5 {
		t.Error("Breadth first search failed.")
	}

	if nodes.Get(0).(TreeNode).Value != 5 {
		t.Error("Breadth first search failed.")
	}

	if nodes.Get(1).(TreeNode).Value != 4 {
		t.Error("Breadth first search failed.")
	}


}
