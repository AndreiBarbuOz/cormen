package subarray

import (
	"testing"
)

func TestFindMaxCrossingSubarrayInvalid(t *testing.T) {
	var intArray []int = make([]int, 1)
	_, _, _, err := findMaxCrossingSubarray(intArray, 1, 0, 0)
	if err == nil {
		t.Errorf("findMaxCrossingSubarray failed to raise error for invalid input")
	}

	intArray[0] = 0
	_, _, _, err = findMaxCrossingSubarray(intArray, 0, 0, 1)
	if err == nil {
		t.Errorf("findMaxCrossingSubarray failed to raise error for array with 1 element")
	}

	intArray = append(intArray, 1)
	_, _, _, err = findMaxCrossingSubarray(intArray, 0, 0, 2)
	if err == nil {
		t.Errorf("findMaxCrossingSubarray failed to raise error for array with low == mid")
	}

	_, _, _, err = findMaxCrossingSubarray(intArray, 0, 1, 2)
	if err == nil {
		t.Errorf("findMaxCrossingSubarray failed to raise error for array with 1 element, high==2")
	}

	intArray = append(intArray, 2)
	_, _, _, err = findMaxCrossingSubarray(intArray, 0, 2, 2)
	if err == nil {
		t.Errorf("findMaxCrossingSubarray failed to raise error for array with mid == high")
	}
}

func TestFindMaxCrossingSubarrayEmpty(t *testing.T) {
	var intArray []int
	low, high, sum, err := findMaxCrossingSubarray(intArray, 0, 0, 0)
	if low != 0 || high != 0 || sum != 0 || err != nil {
		t.Errorf("findMaxCrossingSubarray failed for empty array")
	}
}

func TestFindMaxCrossingSubarrayThreePositive(t *testing.T) {
	intArray := []int{1, 2, 3}
	low, high, sum, err := findMaxCrossingSubarray(intArray, 0, 1, 2)
	if low != 0 || high != 2 || sum != 6 || err != nil {
		t.Errorf("findMaxCrossingSubarray failed for three positive array")
	}

	intArray = []int{1, 2, -1, -3, 5}
	low, high, sum, err = findMaxCrossingSubarray(intArray, 0, 2, 4)
	if low != 0 || high != 4 || sum != 4 || err != nil {
		t.Errorf("findMaxCrossingSubarray failed for mixed five")
	}
}

func TestFindMaxCrossingSubarrayThreeNegative(t *testing.T) {
	intArray := []int{-1, -2, -3}
	low, high, sum, err := findMaxCrossingSubarray(intArray, 0, 1, 2)
	if low != 1 || high != 2 || sum != -5 || err != nil {
		t.Errorf("findMaxCrossingSubarray failed for three negative array")
	}
}

func TestGetMaxSubarrayEmpty(t *testing.T) {
	intArray := []int{}
	low, high, sum, err := getMaxSubarray(intArray, 0, 0)
	if low != 0 || high != 0 || sum != 0 || err != nil {
		t.Errorf("getMaxSubarray failed for empty array")
	}
}

func TestGetMaxSubarrayInvalid(t *testing.T) {
	var intArray []int = make([]int, 1)
	_, _, _, err := getMaxSubarray(intArray, 1, 0)
	if err == nil {
		t.Errorf("getMaxSubarray failed to raise error for invalid input")
	}

	intArray[0] = 0
	_, _, _, err = getMaxSubarray(intArray, 0, 2)
	if err == nil {
		t.Errorf("getMaxSubarray failed to raise error for high greater than array length")
	}

}

func TestGetMaxSubarrayBasic(t *testing.T) {
	intArray := []int{1}
	low, high, sum, err := getMaxSubarray(intArray, 0, 0)
	if low != 0 || high != 0 || sum != 1 || err != nil {
		t.Errorf("getMaxSubarray failed for single value")
	}

	intArray = []int{1, 2}
	low, high, sum, err = getMaxSubarray(intArray, 0, 1)
	if low != 0 || high != 1 || sum != 3 || err != nil {
		t.Errorf("getMaxSubarray failed for two positive values")
	}

	intArray = []int{-1, 2}
	low, high, sum, err = getMaxSubarray(intArray, 0, 1)
	if low != 1 || high != 1 || sum != 2 || err != nil {
		t.Errorf("getMaxSubarray failed for two values, one negative")
	}

	intArray = []int{-1, -2}
	low, high, sum, err = getMaxSubarray(intArray, 0, 1)
	if low != 0 || high != 0 || sum != -1 || err != nil {
		t.Errorf("getMaxSubarray failed for two negative values")
	}

	intArray = []int{1, -2}
	low, high, sum, err = getMaxSubarray(intArray, 0, 1)
	if low != 0 || high != 0 || sum != 1 || err != nil {
		t.Errorf("getMaxSubarray failed for two values, second negative")
	}

}

func TestGetMaxSubarrayCross(t *testing.T) {
	intArray := []int{-2, 3, 2, 5, -4}
	low, high, sum, err := getMaxSubarray(intArray, 0, 4)
	if low != 1 || high != 3 || sum != 10 || err != nil {
		t.Errorf("getMaxSubarray failed for cross midpoint")
	}

	intArray = []int{3, 2, 5}
	low, high, sum, err = getMaxSubarray(intArray, 0, 2)
	if low != 0 || high != 2 || sum != 10 || err != nil {
		t.Errorf("getMaxSubarray failed for cross midpoint, full array")
	}
}

func TestGetMaxSubarrayLeftRight(t *testing.T) {
	intArray := []int{2, 3, 2, -5, -4}
	low, high, sum, err := getMaxSubarray(intArray, 0, 4)
	if low != 0 || high != 2 || sum != 7 || err != nil {
		t.Errorf("getMaxSubarray failed for left subarray")
	}

	intArray = []int{-1, -3, 3, 2, 5}
	low, high, sum, err = getMaxSubarray(intArray, 0, 4)
	if low != 2 || high != 4 || sum != 10 || err != nil {
		t.Errorf("getMaxSubarray failed for right subarray")
	}

	intArray = []int{2, 0, -1, 2, 5}
	low, high, sum, err = getMaxSubarray(intArray, 0, 4)
	if low != 0 || high != 4 || sum != 8 || err != nil {
		t.Errorf("getMaxSubarray failed for full array with 0")
	}

	intArray = []int{1, 0, -1, 0, 5}
	low, high, sum, err = getMaxSubarray(intArray, 0, 4)
	if low != 3 || high != 4 || sum != 5 || err != nil {
		t.Errorf("getMaxSubarray failed for full array with 0")
	}

	intArray = []int{-1, -2, 10, -3, -5}
	low, high, sum, err = getMaxSubarray(intArray, 0, 4)
	if low != 2 || high != 2 || sum != 10 || err != nil {
		t.Errorf("getMaxSubarray failed for full array with large number")
	}

}

func TestMaxSubarrayBruteForceEmpty(t *testing.T) {
	intArray := []int{}
	left, right, err := MaxSubarrayBruteForce(intArray)
	if left != 0 || right != 0 || err != nil {
		t.Errorf("MaxSubarrayBruteForce failed for empty array")
	}
	intArray = append(intArray, 0)
	left, right, err = MaxSubarrayBruteForce(intArray)
	if left != 0 || right != 0 || err != nil {
		t.Errorf("MaxSubarrayBruteForce failed for single item array")
	}
	intArray = []int{1}
	left, right, err = MaxSubarrayBruteForce(intArray)
	if left != 0 || right != 0 || err != nil {
		t.Errorf("MaxSubarrayBruteForce failed for single item array")
	}
}
