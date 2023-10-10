package leetcode

//type Container struct {
//	Width  int
//	Height int
//	Area   int
//}
//
//// O(N^2) is too much
//func MaxAreaN2(height []int) int {
//	var (
//		bestMatch = &Container{0, 0, 0}
//		tmpWidth int
//		tmpHeight int
//		tmpArea int
//	)
//
//	for i := 0; i < len(height); i++ {
//		for j := len(height) - 1; j > 0; j-- {
//			tmpWidth = j - i
//			if height[i] < height[j] {
//				tmpHeight = height[i]
//			} else {
//				tmpHeight = height[j]
//			}
//
//			tmpArea = tmpWidth * tmpHeight
//
//			if tmpArea > bestMatch.Area {
//				bestMatch = &Container{tmpWidth, tmpHeight, tmpArea}
//			}
//		}
//	}
//
//	return bestMatch.Area
//}
//
//func MaxArea(height []int) int {
//
//}
