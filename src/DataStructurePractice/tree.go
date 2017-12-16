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

//用栈进行前序遍历
func stackPreOrder(tree *BinaryTreeNode)  {
	if tree != nil{
		stack:=NewStack()			//新建一个栈
		stack.Push(tree)			//把根结点压入栈
		for !stack.Empty(){
			e:=stack.Pop()				//弹出栈顶元素
			if btNode, ok := e.(*BinaryTreeNode); ok {
				fmt.Println(btNode.data)	//打印当前结点
				if btNode.rightChild != nil{
					stack.Push(btNode.rightChild)	//右结点不为空，则把右节点压入栈
				}
				if btNode.leftChild != nil{
					stack.Push(btNode.leftChild)	//左结点不为空，则把左结点压入栈
				}
			}
		}
	}
}

//用栈实现的中序遍历
func stackMidOrder(tree *BinaryTreeNode)  {
	if tree != nil {
		stack:=NewStack()
		for !stack.Empty() || tree != nil{
			if tree != nil{
				stack.Push(tree)
				tree=tree.leftChild				//指针指向左孩子
			}else {
				//当前元素为空，打印栈顶结点
				e:=stack.Pop()
				if btNode, ok := e.(*BinaryTreeNode); ok {
					fmt.Println(btNode.data)
					tree=btNode.rightChild		//指针指向右孩子
				}
			}
		}
	}
}


func stackPostOrder(tree *BinaryTreeNode){
	if tree != nil {
		s1:=NewStack()
		s2:=NewStack()
		s1.Push(tree)
		for !s1.Empty(){
			e:=s1.Pop()
			if btNode, ok := e.(*BinaryTreeNode); ok {
				tree=btNode
			}
			//s1中弹出的元素全部压入s2
			s2.Push(tree)
			//先压入弹出元素的左孩子
			if tree.leftChild != nil{
				s1.Push(tree.leftChild)
			}
			//后压入弹出元素的右孩子
			if tree.rightChild != nil{
				s1.Push(tree.rightChild)
			}
		}
		//打印s2中的元素
		for !s2.Empty(){
			if btNode, ok := s2.Pop().(*BinaryTreeNode); ok {
				fmt.Println(btNode.data)
			}
		}
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
