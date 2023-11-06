package structures

type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func New[T comparable]() *TreeNode[T] {
	return &TreeNode[T]{}
}

func (t *TreeNode[T]) Insert(_ ...T) {
	panic("not implemented yet")
}

func (t *TreeNode[T]) Remove(_ T) {
	panic("not implemented yet")
}

func (t *TreeNode[T]) Search(_ T) *T {
	panic("not implemented yet")
}

func (t *TreeNode[T]) InOrder(callback func(*TreeNode[T])) {
	inOrderTraversal(t, callback)
}

func inOrderTraversal[T comparable](node *TreeNode[T], callback func(treeNode *TreeNode[T])) {
	if node == nil {
		return
	}

	inOrderTraversal(node.Left, callback)
	callback(node)
	inOrderTraversal(node.Right, callback)
}

func (t *TreeNode[T]) PreOrder(callback func(*TreeNode[T])) {
	preOrderTraversal(t, callback)
}

func preOrderTraversal[T comparable](node *TreeNode[T], callback func(*TreeNode[T])) {
	if node == nil {
		return
	}

	callback(node)
	preOrderTraversal(node.Left, callback)
	preOrderTraversal(node.Right, callback)
}

func (t *TreeNode[T]) PostOrder(callback func(*TreeNode[T])) {
	postOrderTraversal(t, callback)
}

func postOrderTraversal[T comparable](node *TreeNode[T], callback func(*TreeNode[T])) {
	if node == nil {
		return
	}

	postOrderTraversal(node.Left, callback)
	postOrderTraversal(node.Right, callback)
	callback(node)
}

func (t *TreeNode[T]) Depth() int {
	return depth(t)
}

func depth[T comparable](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right)) + 1
}
