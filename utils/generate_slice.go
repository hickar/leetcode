package utils

import (
	"math"
	"math/rand"
	"time"
)

func NewRangeSlice(start, end int) []int {
	newSlice := make([]int, end-start)

	if start < end {
		for i := 0; i < end-start; i++ {
			newSlice[i] = start + i
		}
	} else {
		for i := end; i > start; i-- {
			newSlice[i] = end - i
		}
	}

	return newSlice
}

func NewRandomSlice(length, limit int) []int {
	rand.Seed(time.Now().Unix())
	newSlice := make([]int, length)

	for i := 0; i < length; i++ {
		newSlice[i] = rand.Intn(limit)
	}

	return newSlice
}

func numberOrder(num int) int {
	order := 0
	for num/10 != 0 {
		num /= 10
		order += 1
	}

	return int(math.Floor(math.Pow(float64(10), float64(order))))
}
