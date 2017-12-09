package  main

import (
	"fmt"
)


func main()  {

	testBuildTree()
}



func testBuildTree()  {
	var node BinaryTree
	buildTree(&node)
	fmt.Println("PreOrder:")
	node.preOrderPrint()
	fmt.Println("MidOrder:")
	node.midOrderPrint()
	fmt.Println("PostOrder:")
	node.postOrderPrint()
}

