package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var devCount int
		var devs []int = make([]int, 0)
		fmt.Fscan(in, &devCount)
		for i := 0; i < devCount; i++ {
			var dev int
			fmt.Fscan(in, &dev)
			devs = append(devs, dev)
		}

		devNumbers := make([]int, len(devs))
		for i := range devNumbers {
			devNumbers[i] = i + 1
		}

		for len(devs) > 0 {
			team := []int{devNumbers[0]}
			min := math.MaxInt32
			minIndex := 0
			for i := 1; i < len(devs); i++ {
				diff := int(math.Abs(float64(devs[0] - devs[i])))
				if diff < min {
					min = diff
					minIndex = i
				}
			}
			team = append(team, devNumbers[minIndex])
			fmt.Fprintln(out, team[0], team[1])
			devs = append(devs[:minIndex], devs[minIndex+1:]...)
			devNumbers = append(devNumbers[:minIndex], devNumbers[minIndex+1:]...)
			devs = devs[1:]
			devNumbers = devNumbers[1:]
		}

		fmt.Fprintln(out)
	}
}
