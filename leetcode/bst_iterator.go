package leetcode

type BSTIterator struct {
	cur *TreeNode
	buf []*TreeNode
}

func NewBSTIterator(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	iter.cur, iter.buf = findLeftmost(root, nil)
	return iter
}

func (i *BSTIterator) Next() int {
	if i.cur == nil {
		panic("next element is nil")
	}

	val := i.cur.Val

	switch {
	case i.cur.Right != nil:
		i.cur, i.buf = findLeftmost(i.cur.Right, i.buf)
	case len(i.buf) > 0:
		i.cur = i.buf[len(i.buf)-1]
		i.buf = i.buf[:len(i.buf)-1]
	default:
		i.cur = nil
	}

	return val
}

func (i *BSTIterator) HasNext() bool {
	return len(i.buf) > 0 || i.cur != nil
}

func findLeftmost(root *TreeNode, buf []*TreeNode) (*TreeNode, []*TreeNode) {
	cur := root

	for cur.Left != nil {
		buf = append(buf, cur)
		cur = cur.Left
	}

	return cur, buf
}
