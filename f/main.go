package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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
		var intervals = make([]struct {
			start, end time.Time
		}, timeCount)
		for j := 0; j < timeCount; j++ {
			var line string
			fmt.Fscan(in, &line)
			s := strings.Split(line, "-")
			start, _ := time.Parse("15:04:05", s[0])
			end, _ := time.Parse("15:04:05", s[1])

			intervals[j].start = start
			intervals[j].end = end

			if start.IsZero() || end.IsZero() {
				check = "NO"
			}

			if start.After(end) {
				check = "NO"
			}
		}

		for j := 0; j < len(intervals); j++ {
			if check == "NO" {
				break
			}
			for k := 0; k < len(intervals); k++ {

				if j == k {
					continue
				}

				if ((intervals[j].start.Before(intervals[k].end) || intervals[j].start.Equal(intervals[k].end)) &&
					(intervals[j].start.After(intervals[k].start) || intervals[j].start.Equal(intervals[k].start))) ||
					((intervals[j].end.Before(intervals[k].end) || intervals[j].end.Equal(intervals[k].end)) &&
						(intervals[j].end.After(intervals[k].start) || intervals[j].end.Equal(intervals[k].start))) {
					check = "NO"
					break
				}
			}
			if check == "NO" {
				break
			}

		}
		fmt.Fprintln(out, check)
	}
}
