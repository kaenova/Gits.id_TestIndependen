/*
Gits.ID Assessment
Kaenova Mahendra Auditama
kaenova@gmail.com

Note: Pada kasus ini, saya tidak mempertimbangkan adanya tanda baca
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palindrom(input string) bool{
	var idx_r, idx_l, max, i int

	input = strings.Replace(input, " ", "", -1)
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	idx_r = len(input)-1
	idx_l = 0
	max = idx_r / 2

	for i = 0; i <= max ; i++ {
		if input[idx_l + i] != input[idx_r - i]{
			return false
		}
	} 
	return true
}

func main(){
	var input string
	fmt.Println("Masukkan Input: ")

	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Palindrom:", palindrom(input))
}