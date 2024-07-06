package numberUtil

import (
	"math"
)

// Min 返回整数中的最小值
func Min(intList []int) int {
	if len(intList) == 0 {
		panic("intList is empty")
	}
	minVal := intList[0]
	for _, v := range intList {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// Max 返回两个整数中的最大值
func Max(intArr []int) int {
	if len(intArr) == 0 {
		panic("intArr is empty")
	}
	maxVal := intArr[0]
	for _, v := range intArr {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func ScaleFloat64(number float64, scale int) float64 {
	result := math.Round(number*100) / 100
	return result
}
