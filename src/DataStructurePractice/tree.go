package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	data 		string				//数据
	leftChild 	*BinaryTreeNode		//左孩子
	rightChild	*BinaryTreeNode		//右孩子
}


//面向对象初始化二叉树
func (node *BinaryTreeNode)initTree(){
	var input string
	fmt.Println("Please input the data of node")
	fmt.Scanln(&input)
	if input=="#"{
		node=nil
	}else
	{
		node.data=input
		fmt.Println("The data you have input :"+ node.data)
		node.leftChild=new(BinaryTreeNode)
		node.leftChild.initTree()
		node.rightChild=new(BinaryTreeNode)
		node.rightChild.initTree()
	}
}


//前序遍历输出结点
func preOrder(tree  *BinaryTreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.data)
	preOrder(tree.leftChild)
	preOrder(tree.rightChild)

}

//中序遍历输出结点
func midOrder(tree  *BinaryTreeNode) {
	if tree == nil {
		return
	}
	midOrder(tree.leftChild)
	fmt.Println(tree.data)
	midOrder(tree.rightChild)
}

//后序遍历输出结点
func postOrder(tree  *BinaryTreeNode) {
	if tree == nil {
		return
	}
	postOrder(tree.leftChild)
	postOrder(tree.rightChild)
	fmt.Println(tree.data)
}

//面向对象前序遍历
func (tree *BinaryTreeNode)preOrderPrint() {
	if tree == nil {
		return
	}
	fmt.Println(tree.data)
	tree.leftChild.preOrderPrint()
	tree.rightChild.preOrderPrint()
}


//面向对象中序遍历
func (tree *BinaryTreeNode)midOrderPrint() {
	if tree == nil {
		return
	}
	tree.leftChild.midOrderPrint()
	fmt.Println(tree.data)
	tree.rightChild.midOrderPrint()
}


//面向对象后序遍历
func (tree *BinaryTreeNode)postOrderPrint() {
	if tree == nil {
		return
	}
	tree.leftChild.postOrderPrint()
	tree.rightChild.postOrderPrint()
	fmt.Println(tree.data)
}





//新建二叉树
func buildTree(node *BinaryTreeNode){
	var input string
	fmt.Println("Please input your something: ")
	fmt.Scanln(&input)
	if input == "#"{
		node=nil		//结点为空，函数返回
	}else
	{
		node.data=input
		fmt.Println("You input:"+node.data)
		node.leftChild=new(BinaryTreeNode)
		buildTree(node.leftChild)
		node.rightChild=new(BinaryTreeNode)
		buildTree(node.rightChild)
	}
}
