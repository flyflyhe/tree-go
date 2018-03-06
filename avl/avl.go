package main

import (
)

type AvlTree struct {

	Key int
	
	Height int

	Used bool

	Left *AvlTree

	Right *AvlTree
}

//左左情况  右旋 top顶级节点
func RotateLL(top *AvlTree) *AvlTree {
	
	//newTop: 需要作为顶级节点的元素
	newTop := top.Left

	//先截断当前节点的左孩子
	top.Left = nil

	//将当前top节点作为newTop的右孩子
	newTop.Right = top

	//计算当前两个节点的高度
	top.Height = top.MaxHeight()
	newTop.Height = newTop.MaxHeight()

	return newTop
}

//右右情况 左旋
func RotateRR(top *AvlTree) *AvlTree {

	newTop := top.Right

	top.Right = nil

	newTop.Left = top

	top.Height = top.MaxHeight()
	newTop.Height = newTop.MaxHeight()

	return top
}

//左右情况 top是当前顶级节点
func RotateLR(top *AvlTree) *AvlTree {

	//先进行RR情况旋转
	top.Left = RotateRR(top.Left)

	//在进行LL情况旋转

	return RotateLL(top)
}

//右左情况 
func RotateRL(top *AvlTree) *AvlTree {

	//先进行左左情况旋转
	top.Right = RotateLL(top.Right)

	//再执行右右情况旋转
	return RotateRR(top)
}

//根据当前节点计算节点的最大高度
func (node *AvlTree) MaxHeight() int {
	leftHeight := 0
	rightHeight := 0
	left := node.Left
	right := node.Right

	//左节点查找
	for {
		if left == nil {
			break
		}
		left = left.Left
		leftHeight += 1
	}

	//右节点查找
	for {
		if right == nil {
			break
		}
		right = right.Right
		rightHeight += 1
	}

	if leftHeight >= rightHeight {
		return leftHeight
	}

	return rightHeight
}

func main () {

}