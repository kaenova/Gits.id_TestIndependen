/*
Gits.ID Assessment
Kaenova Mahendra Auditama
kaenova@gmail.com
*/

package main

import "fmt"

func main(){
	var input int

	fmt.Print("Masukkan Input: ")
	fmt.Scan(&input)

	if input % 5 == 0 && input % 3 ==0 {
		fmt.Println("Hello World")
	} else if input % 3 == 0 {
		fmt.Print("Hello")
	} else if input % 5 == 0 {
		fmt.Println("World")
	}
}
