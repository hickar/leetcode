package leetcode

import "math"

/*
134. Gas Station (Medium)
There are n gas stations along a circular route,
where the amount of gas at the ith station is gas[i].
You have a car with an unlimited gas tank and it costs cost[i] of gas
to travel from the ith station to its next (i + 1)th station.
You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost,
return the starting gas station's index if you can travel around the circuit once in the clockwise direction,
otherwise return -1.
If there exists a solution, it is guaranteed to be unique


Example 1:
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.

Example 2:
Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.
*/

func canCompleteCircuitOptimized(gas []int, cost []int) int {
	n := len(gas)
	f, g, start := 0, 0, 0

	for i := 0; i < n; i++ {
		g += gas[i] - cost[i]
		f += gas[i] - cost[i]
		if f < 0 {
			start = i + 1
			f = 0
		}
	}

	if g < 0 {
		return -1
	}

	return start
}

func canCompleteCircuit(gas []int, cost []int) int {
	var (
		ok       bool
		nextPos  int
		startPos = findStartPos(gas, cost)
	)

	for i, r := startPos, 0; r < len(gas); r++ {
		if i >= len(gas) {
			i = 0
		}

		nextPos, ok = travelFromPos(gas, cost, i)
		if ok {
			return i
		}
		if i < startPos && nextPos >= startPos {
			break
		}
		if i > startPos && nextPos == startPos {
			break
		}
		if i == nextPos {
			i++
			continue
		}
		if i >= startPos {
			i = max(nextPos, i+1)
			continue
		}
		if i < startPos && nextPos <= startPos {
			i = max(nextPos, i+1)
			continue
		}
	}

	return -1
}

func travelFromPos(gas, cost []int, start int) (int, bool) {
	var (
		fuel       int
		travelCost int
	)

	for i, r := start, 0; r < len(gas); r++ {
		if i == len(gas) {
			i = 0
		}

		fuel += gas[i]
		travelCost = cost[i]

		if i == len(gas)-1 {
			i = -1
		}

		if fuel < travelCost || fuel <= 0 {
			return i + 1, false
		}

		fuel -= travelCost
		i++
	}

	return start, true
}

func findStartPos(gas, cost []int) int {
	var (
		minFuel = math.MaxInt64
		pos     = -1
	)

	fuel := 0
	for i := 0; i < len(gas); i++ {
		fuel += gas[i]
		fuel -= cost[i]

		if fuel < minFuel {
			pos = i
			minFuel = fuel
		}
	}

	pos += 1
	if pos == len(gas) {
		return 0
	}

	return pos
}
