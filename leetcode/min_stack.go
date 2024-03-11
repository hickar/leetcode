package leetcode

import "math"

/*
155. Min Stack (Medium)
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
Implement the MinStack class:
- MinStack() initializes the stack object.
- void push(int val) pushes the element val onto the stack.
- void pop() removes the element on the top of the stack.
- int top() gets the top element of the stack.
- int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.


Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:
- -2^31 <= val <= 2^31 - 1
- Methods pop, top and getMin operations will always be called on non-empty stacks.
- At most 3 * 104 calls will be made to push, pop, top, and getMin.
*/

type msNode struct {
	Value   int
	PrevMin int
}

type MinStack struct {
	items []*msNode
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	node := &msNode{
		Value:   val,
		PrevMin: math.MaxInt64,
	}
	if len(s.items) > 0 {
		prevNode := s.items[len(s.items)-1]
		node.PrevMin = min(prevNode.Value, prevNode.PrevMin)
	}

	s.items = append(s.items, node)
}

func (s *MinStack) Pop() {
	if len(s.items) == 0 {
		panic("trying to pop out of empty stack")
	}

	s.items = s.items[:len(s.items)-1]
}

func (s *MinStack) Top() int {
	if len(s.items) == 0 {
		panic("no elements in stack")
	}

	return s.items[len(s.items)-1].Value
}

func (s *MinStack) GetMin() int {
	if len(s.items) == 0 {
		panic("trying to get min of empty stack")
	}

	topNode := s.items[len(s.items)-1]
	return min(topNode.Value, topNode.PrevMin)
}
