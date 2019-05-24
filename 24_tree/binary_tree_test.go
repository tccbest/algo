package _4_tree

import "testing"

func TestNode_PreOrder(t *testing.T) {
	a := &Node{1, nil, nil}
	b := &Node{2, nil, nil}
	c := &Node{3, nil, nil}
	d := &Node{4, nil, nil}
	e := &Node{5, nil, nil}
	a.left = b
	a.right = c
	b.left = d
	b.right = e

	a.PreOrder()
}

func TestNode_InOrder(t *testing.T) {
	a := &Node{1, nil, nil}
	b := &Node{2, nil, nil}
	c := &Node{3, nil, nil}
	d := &Node{4, nil, nil}
	e := &Node{5, nil, nil}
	a.left = b
	a.right = c
	b.left = d
	b.right = e

	a.InOrder()
}

func TestNode_PostOrder(t *testing.T) {
	a := &Node{1, nil, nil}
	b := &Node{2, nil, nil}
	c := &Node{3, nil, nil}
	d := &Node{4, nil, nil}
	e := &Node{5, nil, nil}
	a.left = b
	a.right = c
	b.left = d
	b.right = e

	a.PostOrder()
}
