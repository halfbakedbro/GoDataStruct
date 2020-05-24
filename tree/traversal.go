package tree

import "../queue"

/**
* LevelOrder Traversal
* Inorder Traversal
* PreOrder Traversal
* PostOrder Traversal
* Vertical Order Traversal
* Diagnol Traversal
* Boundary Traversal
* Spiral Level Order Traversal
**/

func LevelOrder(tree *BinTree) {
	if tree == nil || tree.root == nil {
		return
	}

	qu := queue.NewQueue()

	qu.Enqueue(tree.root)

	for !qu.IsEmpty() {
		node := qu.Dequeue()
		//print node val
		// push left and right
		if node.left != nil {
			queue.Enqueue(node.left)
		}

		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}

}
