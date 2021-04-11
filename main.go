package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)
import "./subarray"


func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
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
	file, err := os.Open("input1.txt")
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
	subarray.GetMaxSubarray(integers)
	fmt.Printf("%v\n", integers)
}
