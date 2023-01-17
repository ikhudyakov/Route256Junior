package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var days int
		var temp []int = make([]int, 0)
		var reports map[int]int = make(map[int]int)
		var check = "YES"
		fmt.Fscan(in, &days)
		for j := 0; j < days; j++ {
			var report int
			fmt.Fscan(in, &report)
			temp = append(temp, report)

			if j == 0 {
				reports[report] = j
			}
			if j > 0 && temp[j-1] != report {
				if _, ok := reports[report]; !ok {
					reports[report] = j
				} else {
					check = "NO"
				}
			}
		}
		fmt.Fprintln(out, check)
	}
}
