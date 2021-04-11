package subarray

import (
	"fmt"
	"math"
)

func GetMaxSubarray(arr []int) (int, int) {
	return 0, 0
}

func findMaxCrossingSubarray(arr []int, low int, mid int, high int) (maxLeft int, maxRight int, maxSum int, err error) {
	if len(arr) == 0 {
		return 0, 0, 0, nil
	}

	if low >= mid || mid >= high {
		return 0, 0, 0, fmt.Errorf("invalid input values")
	}

	if low < 0 || high >= len(arr) {
		return 0, 0, 0, fmt.Errorf("index out of bounds")
	}

	var sumLeft int = math.MinInt64
	maxLeft = mid
	var sum int = 0
	for i := mid; i >= low; i-- {
		sum += arr[i]
		if sum > sumLeft {
			sumLeft = sum
			maxLeft = i
		}
	}

	var sumRight int = math.MinInt64
	maxRight = mid + 1
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += arr[i]
		if sum > sumRight {
			sumRight = sum
			maxRight = i
		}
	}

	maxSum = sumLeft + sumRight

	return maxLeft, maxRight, maxSum, nil
}
