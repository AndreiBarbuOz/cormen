package subarray

import (
	"fmt"
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

	intArray = append(intArray, 2)
	_, _, _, err = findMaxCrossingSubarray(intArray, 0, 2, 2)
	if err == nil {
		t.Errorf("findMaxCrossingSubarray failed to raise error for array with mid == high")
	}
	fmt.Printf("array %v\n", intArray)

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
	if low != 0 || high != 0 || sum != 3 || err != nil {
		t.Errorf("findMaxCrossingSubarray failed for three positive array")
	}

}
