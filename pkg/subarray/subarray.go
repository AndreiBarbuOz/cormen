package subarray

import (
	"fmt"
	"math"
)

func MaxSubarray(arr []int) (left int, right int, err error) {
	return 0, 0, nil
}

func getMaxSubarray(arr []int, low int, high int) (maxLeft int, maxRight int, maxSum int, err error) {
	if len(arr) == 0 {
		return 0, 0, 0, nil
	}
	if low < 0 || high >= len(arr) {
		return 0, 0, 0, fmt.Errorf("index out of bounds")
	}
	if low > high {
		return 0, 0, 0, fmt.Errorf("invalid input values")
	}
	if low == high {
		return low, low, arr[low], nil
	}
	if low == high-1 { // 2 elements in array
		if arr[low] >= 0 && arr[high] >= 0 {
			return low, high, arr[low] + arr[high], nil
		} else if arr[low] < arr[high] {
			return high, high, arr[high], nil
		} else {
			return low, low, arr[low], nil
		}
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum, err := getMaxSubarray(arr, low, mid)
	rightLow, rightHigh, rightSum, err := getMaxSubarray(arr, mid+1, high)
	crossLow, crossHigh, crossSum, err := findMaxCrossingSubarray(arr, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum, nil
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum, nil
	} else {
		return crossLow, crossHigh, crossSum, nil
	}
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
