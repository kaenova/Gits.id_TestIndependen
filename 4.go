/*
Gits.ID Assessment
Kaenova Mahendra Auditama
kaenova@gmail.com
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(input string) string{
	var rev_string string
	for i:= len(input)-1 ; i > -1; i--{
		rev_string = rev_string + string(input[i])
	}
	return rev_string
}

func main(){
	var input string
	fmt.Print("Masukkan Input: ")
	
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("String Reverse:", reverseString(input))
}