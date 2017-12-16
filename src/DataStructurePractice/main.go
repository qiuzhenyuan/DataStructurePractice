package  main

import (
	"fmt"
)


func main()  {
	//fmt.Println("hi")
	//var bb interface{}
	//jj:=bb.(*string)
	//
	//fmt.Println(jj)
testStack()
}



func testBuildTree()  {
	var node BinaryTreeNode
	buildTree(&node)
	fmt.Println("PreOrder:")
	node.preOrderPrint()
	fmt.Println("MidOrder:")
	node.midOrderPrint()
	fmt.Println("PostOrder:")
	node.postOrderPrint()
}

func testInitTree(){
	var node BinaryTreeNode
	buildTree(&node)
	fmt.Println("PreOrder:")
	preOrder(&node)
	fmt.Println("MidOrder:")
	midOrder(&node)
	fmt.Println("PostOrder:")
	postOrder(&node)

}

func testStack()  {
	var node BinaryTreeNode
	buildTree(&node)
	//fmt.Println("PreOrder:")
	//preOrder(&node)
	//fmt.Println("StackPreOrder:")
	//stackPreOrder(&node)
	//
	fmt.Println("MidOrder:")
	midOrder(&node)
	fmt.Println("StackMidOrder:")
	stackMidOrder(&node)

	fmt.Println("StackPostOrder:")
	stackPostOrder(&node)

}
