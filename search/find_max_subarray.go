package search

import "math"

func FindMaxSubarray(arr []int, low, high int) (int, int, int) {
	if low == high {
		return arr[low], low, high
	} else {
		mid := int(math.Floor(float64(low+high) / 2))
		leftSum, leftLow, leftHigh := FindMaxSubarray(arr, low, mid)
		rightSum, rightLow, rightHigh := FindMaxSubarray(arr, mid+1, high)
		crossSum, crossLow, crossHigh := FindCrossMaxSubarray(arr, low, mid, high)

		if leftSum > rightSum && leftSum > crossSum {
			return leftSum, leftLow, leftHigh
		} else if rightSum > leftSum && rightSum > crossSum {
			return rightSum, rightLow, rightHigh
		} else {
			return crossSum, crossLow, crossHigh
		}
	}
}

func FindCrossMaxSubarray(arr []int, low, mid, high int) (int, int, int) {
	sum := 0
	leftMax, rightMax := 0, 0
	leftSum, rightSum := math.MinInt32, math.MinInt32

	for i := mid; i > low; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
			leftMax = i
		}
	}

	sum = 0
	for j := mid + 1; j < high; j++ {
		sum += arr[j]
		if sum > rightSum {
			rightSum = sum
			rightMax = j
		}
	}

	return leftSum + rightSum, leftMax, rightMax
}
