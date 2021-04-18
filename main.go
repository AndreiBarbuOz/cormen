package main

import (
	"bufio"
	"fmt"
	"github.com/AndreiBarbuOz/cormen/pkg/subarray"
	"io"
	"os"
	"strconv"
	"time"
)

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	_, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}
	var ret []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return ret, err
		}
		ret = append(ret, x)
	}
	return ret, scanner.Err()
}

func main() {
	for i := 1; i < 25; i++ {
		file, err := os.Open(fmt.Sprintf("input%d.txt", i))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := io.Reader(file)
		integers, err := readInts(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		start := time.Now()
		diLeft, diRight, err := subarray.MaxSubarray(integers)
		diDuration := time.Since(start)
		start = time.Now()
		bfLeft, bfRight, err := subarray.MaxSubarrayBruteForce(integers)
		bfDuration := time.Since(start)
		fmt.Printf("i = %d len = %v div et imp=%v brute force=%v\n", i, len(integers), diDuration, bfDuration)
		if bfLeft != diLeft || bfRight != diRight {
			fmt.Printf("responses don't match\n")
		}
	}
}
