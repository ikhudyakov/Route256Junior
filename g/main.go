package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var userCount, pairCount int
	fmt.Fscan(in, &userCount, &pairCount)

	users := make(map[int]map[int]bool)

	for i := 0; i < pairCount; i++ {
		var user int
		var friend int
		fmt.Fscan(in, &user, &friend)

		_, ok := users[user]
		if !ok {
			users[user] = make(map[int]bool)
			users[user][friend] = false
		} else {
			users[user][friend] = false
		}

		_, ok = users[friend]
		if !ok {
			users[friend] = make(map[int]bool)
			users[friend][user] = false
		} else {
			users[friend][user] = false
		}
	}

	for i := 1; i < userCount+2; i++ {
		var tempArr []int
		tempMap := make(map[int]int)
		for k, _ := range users[i] {
			for m, _ := range users[k] {
				if m == i {
					continue
				}
				_, ok := users[i][m]
				if ok {
					continue
				}
				_, ok = users[i][k]
				if !ok {
					continue
				} else {
					tempMap[m]++
				}
			}
		}

		maxFriends := 0
		for _, val := range tempMap {
			if val > maxFriends {
				maxFriends = val
			}
		}

		for key, val := range tempMap {
			if val == maxFriends {
				tempArr = append(tempArr, key)
			}
		}

		sort.Ints(tempArr)
		if i == userCount+1 {
			break
		}

		if len(tempArr) == 0 {
			fmt.Fprintln(out, "0")
			continue
		} else {
			fmt.Fprintln(out, strings.Trim(fmt.Sprint(tempArr), "[]"))
		}
	}
}
