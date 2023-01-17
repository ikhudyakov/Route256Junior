package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	id                int
	friendsId         []int
	possibleFriendsId []int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var userCount, pairCount int
	fmt.Fscan(in, &userCount, &pairCount)

	users := make([]User, userCount)
	for i := 0; i < userCount; i++ {
		user := User{
			id:                i + 1,
			friendsId:         make([]int, 0),
			possibleFriendsId: make([]int, 0),
		}
		users[i] = user
	}

	for i := 0; i < pairCount; i++ {
		var user, friend int
		fmt.Fscan(in, &user, &friend)
		users[user-1].friendsId = append(users[user-1].friendsId, friend)
		users[friend-1].friendsId = append(users[friend-1].friendsId, user)
	}

	for _, u := range users {
		u.possibleFriendsId = findCommonFriends(u.id, users)
		fmt.Fprintln(out, u.possibleFriendsId)
	}
}

func findCommonFriends(user int, users []User) []int {
	var commonFriends []int
	for _, u := range users {
		if u.id == user || isFriend(user, u.id, users) {
			continue
		}
		for _, f := range u.friendsId {
			if isFriend(user, f, users) {
				commonFriends = append(commonFriends, u.id)
				break
			}
		}
	}
	commonFriends = removeDuplicates(commonFriends)
	return commonFriends
}

func isFriend(u1, u2 int, users []User) bool {
	for _, f := range users[u1-1].friendsId {
		if f == u2 {
			return true
		}
	}
	return false
}

func removeDuplicates(s []int) []int {
	seen := make(map[int]bool)
	j := 0
	for i, x := range s {
		if _, ok := seen[x]; !ok {
			seen[x] = true
			s[j] = s[i]
			j++
		}
	}
	return s[:j]
}
