package  main

import (
	"fmt"
)


func main()  {

	testInitTree()
	testBuildTree()
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
