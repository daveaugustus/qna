// Written by: Dave
// Write a function to find consecutive repeating characters in a string
// put them in a slice and return the slice of string
// *****************************************************************
// eg:
// input string := "aaabccdddefgaakllllkkkkll"
// output slice := ["aa", "cc", "ddd", "aa", "llll", "kkkk", "ll"]
//
//*******************************************************************

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "aaabccdddefgaakllllkkkkll"
	fmt.Println(ConsecutiveStrToSlice(str1))
}

func ConsecutiveStrToSlice(str string) []string {
	var i int
	if len(str) <= 1 {
		return nil
	}
	list := make([]string, 0, 0)
	occurrence := 0

	for i = 0; i < len(str)-1; i++ {
		if str[i+1] == str[i] {
			occurrence++
		} else {
			if occurrence > 0 {
				list = append(list, strings.Repeat(string(str[i]), occurrence+1))
			}
			occurrence = 0
		}
	}
	if occurrence > 0 {
		list = append(list, strings.Repeat(string(str[i]), occurrence+1))
	}
	return list
}
