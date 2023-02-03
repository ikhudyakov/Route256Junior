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

	var size int
	fmt.Fscan(in, &size)

	rymMap := make(map[string][]string)
	var anyWord [2]string

	for i := 0; i < size; i++ {
		var word string
		fmt.Fscan(in, &word)
		if i < 2 {
			anyWord[i] = word
		}
		origWord := word
		for j := len(word) - 1; j >= 0; j-- {
			_, ok := rymMap[word[j:]]
			if !ok {
				var libArr []string
				libArr = append(libArr, origWord)
				rymMap[word[j:]] = libArr
			} else {
				tempLib := rymMap[word[j:]]
				tempLib = append(tempLib, origWord)
				rymMap[word[j:]] = tempLib
			}
		}
	}
	var checkSize int
	fmt.Fscan(in, &checkSize)
	var result string
	for k := 0; k < checkSize; k++ {
		var word string
		fmt.Fscan(in, &word)
		result = anyWord[1]
		for l := len(word) - 1; l >= 0; l-- {
			val, ok := rymMap[word[l:]]
			if ok {
				for m := 0; m < len(val); m++ {
					if word != val[m] {
						result = val[m]
						break
					}
				}
			}
		}

		fmt.Fprintln(out, result)
	}
}
