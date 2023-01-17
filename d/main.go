package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Table [][]int

func (t Table) Len() int {
	return len(t)
}

func (t Table) Less(i, j int) bool {
	return t[i][sortColumn] < t[j][sortColumn]
}

func (t Table) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

var sortColumn int

func sortTable(table Table, column int) {
	sortColumn = column
	sort.SliceStable(table, func(i, j int) bool {
		return table[i][column] < table[j][column]
	})
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var x, y, columnCount int
		fmt.Fscan(in, &y, &x)
		var table Table = make(Table, y)
		for i := range table {
			table[i] = make([]int, x)
			for j := 0; j < x; j++ {
				var a int
				fmt.Fscan(in, &a)
				table[i][j] = a
			}
		}
		fmt.Fscan(in, &columnCount)

		for i := 0; i < columnCount; i++ {
			var column int
			fmt.Fscan(in, &column)
			sortTable(table, column-1)
		}

		for _, v := range table {
			for _, v1 := range v {
				fmt.Fprint(out, v1, " ")
			}
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
}
