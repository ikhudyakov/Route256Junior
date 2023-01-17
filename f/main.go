package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval []struct {
	start, end int
}

func (i Interval) Len() int {
	return len(i)
}

func (i Interval) Less(k, j int) bool {
	return i[k].start < i[j].start
}

func (i Interval) Swap(k, j int) {
	i[k], i[j] = i[j], i[k]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var timeCount int
		var check = "YES"
		fmt.Fscan(in, &timeCount)
		var intervals = make(Interval, timeCount)
		for j := 0; j < timeCount; j++ {
			var line string
			fmt.Fscan(in, &line)
			s := strings.Split(line, "-")
			start := validate(s[0])
			end := validate(s[1])

			intervals[j].start = start
			intervals[j].end = end

			if start < 0 || end < 0 || start > end {
				check = "NO"
			}
		}

		sort.Sort(intervals)
		for i := 0; i < intervals.Len()-1; i++ {
			if intervals[i].end >= intervals[i+1].start {
				check = "NO"
				break
			}
		}

		fmt.Fprintln(out, check)
	}
}

func validate(line string) int {
	arr := strings.Split(line, ":")
	h, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	s, _ := strconv.Atoi(arr[2])
	if h < 0 || h > 23 || m < 0 || m > 59 || s < 0 || s > 59 {
		return -1
	}
	return s + m*60 + h*3600
}
