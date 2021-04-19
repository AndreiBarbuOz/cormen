package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"math/rand"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{Use: "cormen"}
var filePath string

var subarrayCmd = &cobra.Command{
	Use:   "subarray",
	Short: "subarray manipulation",
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate test data for subarray testing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generate file %v\n", filePath)
	},
}

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute subarray against test data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Execute\n")
	},
}

func Execute(args []string) {

	executeCmd.PersistentFlags().StringVar(&filePath, "input", "input.txt", "Input test file")
	generateCmd.PersistentFlags().StringVar(&filePath, "output", "output.txt", "Generated test file")

	subarrayCmd.AddCommand(generateCmd)
	subarrayCmd.AddCommand(executeCmd)
	rootCmd.AddCommand(subarrayCmd)
	if err := subarrayCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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

func generateInts(size int) *[]int {
	var ret []int
	ret = make([]int, size)
	for i := 0; i < size; i++ {
		ret[i] = rand.Intn(1000) - 500
	}
	return &ret
}

func generateFile(filePath string, size int) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("generateFile opening file %s failed: %v", filePath, err)
	}
	if size < 0 {
		panic("size < 0")
	}
	var values []int = *generateInts(size)

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(strconv.Itoa(size) + "\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, value := range values {
		_, err := writer.WriteString(strconv.Itoa(value) + "\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	writer.Flush()
	file.Close()

	return nil
}

//func Execute() {
//	for i := 1; i < 25; i++ {
//		file, err := os.Open(fmt.Sprintf("input%d.txt", i))
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//		reader := io.Reader(file)
//		integers, err := readInts(reader)
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//		start := time.Now()
//		diLeft, diRight, err := subarray.MaxSubarray(integers)
//		diDuration := time.Since(start)
//		start = time.Now()
//		bfLeft, bfRight, err := subarray.MaxSubarrayBruteForce(integers)
//		bfDuration := time.Since(start)
//		fmt.Printf("i = %d len = %v div et imp=%v brute force=%v\n", i, len(integers), diDuration, bfDuration)
//		if bfLeft != diLeft || bfRight != diRight {
//			fmt.Printf("responses don't match\n")
//		}
//	}
//}
