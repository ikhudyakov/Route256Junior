package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameMap [][]rune

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)
	var x, y int

	for i := 0; i < testCount; i++ {
		fmt.Fscan(in, &x)
		fmt.Fscanln(in, &y)
		rowHexArr := make(map[int][]rune)
		objectsMap := make(map[string]bool)

		for i := 0; i < x; i++ {
			str, _ := in.ReadString(10)
			str = strings.Trim(str, "\n")

			rowHex := []rune(str)
			rowHex = rowHex[0 : len(rowHex)-1]
			for _, val := range rowHex {
				if val == 46 {
					continue
				}

				objectsMap[string(val)] = false
			}
			rowHexArr[i] = rowHex
		}

		var isError bool
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				if rowHexArr[i][j] == 46 {
					continue
				}
				chkdVal := rowHexArr[i][j]
				if objectsMap[string(chkdVal)] {
					isError = true
					break
				}
				var ValToCheckArr [][12]int
				coordinateArr := [12]int{i - 1, j - 1, i - 1, j + 1, i, j + 2, i + 1, j + 1, i + 1, j - 1, i, j - 2}
				rowHexArr[i][j] = 46
				ValToCheckArr = append(ValToCheckArr, coordinateArr)

				traveler := func(chkArr [][12]int) [][12]int {
					var newchkArr [][12]int
					for i := 0; i < len(chkArr); i++ {
						findObjects := 6
						crdnArr := chkArr[i]
						for j := 0; j < 12; j = j + 2 {
							tR := crdnArr[j]
							tC := crdnArr[j+1]
							switch {
							case tR < 0 || tC < 0:
								findObjects--
							case tR >= x || tC >= y:
								findObjects--
							case rowHexArr[tR][tC] == 46:
								findObjects--
							case rowHexArr[tR][tC] == chkdVal:
								findObjects--
								rowHexArr[tR][tC] = 46
								newcrdnArr := [12]int{tR - 1, tC - 1, tR - 1, tC + 1, tR, tC + 2, tR + 1, tC + 1, tR + 1, tC - 1, tR, tC - 2}
								newchkArr = append(newchkArr, newcrdnArr)
							default:
								findObjects--
							}
						}
					}
					if len(newchkArr) == 0 {
						objectsMap[string(chkdVal)] = true
					}
					return newchkArr
				}
				recur(traveler, ValToCheckArr)
			}
			if isError {
				break
			}
		}
		if !isError {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
func recur(f func([][12]int) [][12]int, arr [][12]int) (func([][12]int) [][12]int, [][12]int) {
	if len(arr) == 0 {
		return nil, nil
	}
	res := f(arr)
	return recur(f, res)
}
