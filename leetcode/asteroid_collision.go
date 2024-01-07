package leetcode

/*
735. Asteroid Collision (Medium)

We are given an array asteroids of integers representing asteroids in a row.
For each asteroid, the absolute value represents its size,
and the sign represents its direction (positive meaning right, negative meaning left).
Each asteroid moves at the same speed.
Find out the state of the asteroids after all collisions.
If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode.
Two asteroids moving in the same direction will never meet.


Example 1:
Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

Example 2:
Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.

Example 3:
Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
*/

type AsteroidStack struct {
	items []int
	size  int
}

func NewStack(size int) *AsteroidStack {
	return &AsteroidStack{
		items: make([]int, 0, size),
	}
}

func (s *AsteroidStack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *AsteroidStack) Push(item int) {
	s.items = append(s.items, item)
	s.size++
}

func (s *AsteroidStack) Pop() int {
	n := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.size--
	return n
}

func (s *AsteroidStack) Size() int {
	return s.size
}

func (s *AsteroidStack) Slice() []int {
	scopy := make([]int, len(s.items))
	copy(scopy, s.items)
	return scopy
}

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) < 2 {
		return asteroids
	}

	stack := NewStack(len(asteroids))
	stack.Push(asteroids[0])

	var prev int
	var cur int
	for i := 1; i < len(asteroids); i++ {
		cur = asteroids[i]
		if stack.Size() == 0 {
			stack.Push(cur)
			continue
		}

		prev = stack.Peek()
		if !shouldCollide(prev, cur) {
			stack.Push(cur)
			continue
		}

		simulateCollisions(stack, cur)
	}

	return stack.Slice()
}

func simulateCollisions(stack *AsteroidStack, cur int) {
	var (
		prev = stack.Peek()
		diff int
	)

	for {
		if stack.Size() == 0 {
			stack.Push(cur)
			return
		}

		prev = stack.Peek()
		if !shouldCollide(prev, cur) {
			stack.Push(cur)
			return
		}

		diff = abs(prev) - abs(cur)
		switch {
		case diff > 0:
			return
		case diff < 0:
			stack.Pop()
		case diff == 0:
			stack.Pop()
			return
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func shouldCollide(n1, n2 int) bool {
	return n1 > -1 && n2 < 0
}
