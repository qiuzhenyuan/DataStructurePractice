package main

import (
	"fmt"
)

type BinaryTree struct {
	data 		string			//
	leftChild 	*BinaryTree		//
	rightChild	*BinaryTree		//
}



//新建二叉树
func buildTree(node *BinaryTree){
	var input string
	fmt.Println("Please input your something: ")
	fmt.Scanln(&input)
	if input == "#"{
		node=nil		//结点为空，函数返回
	}else
	{
		node.data=input
		fmt.Println("You input:"+node.data)
		node.leftChild=new(BinaryTree)
		buildTree(node.leftChild)
		node.rightChild=new(BinaryTree)
		buildTree(node.rightChild)
	}
}


//前序遍历输出结点
func preOrder(tree  *BinaryTree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.data)
	preOrder(tree.leftChild)
	preOrder(tree.rightChild)

}



//中序遍历输出结点
func midOrder(tree  *BinaryTree) {
	if tree == nil {
		return
	}
	preOrder(tree.leftChild)
	fmt.Println(tree.data)
	preOrder(tree.rightChild)
}

//后序遍历输出结点
func postOrder(tree  *BinaryTree) {
	if tree == nil {
		return
	}
	preOrder(tree.leftChild)
	preOrder(tree.rightChild)
	fmt.Println(tree.data)
}

//面向对象前序遍历
func (tree *BinaryTree)preOrderPrint() {
	if tree == nil {
		return
	}
	fmt.Println(tree.data)
	tree.leftChild.preOrderPrint()
	tree.rightChild.preOrderPrint()
}

//面向对象中序遍历
func (tree *BinaryTree)midOrderPrint() {
	if tree == nil {
		return
	}
	fmt.Println(tree.data)
	tree.leftChild.midOrderPrint()
	tree.rightChild.midOrderPrint()
}


//面向对象后序遍历
func (tree *BinaryTree)postOrderPrint() {
	if tree == nil {
		return
	}
	fmt.Println(tree.data)
	tree.leftChild.postOrderPrint()
	tree.rightChild.postOrderPrint()
}
