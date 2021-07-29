/*
Gits.ID Assessment
Kaenova Mahendra Auditama
kaenova@gmail.com
*/

package main

import (
	"fmt"
	"strings"
)

func emailCheck(input string) bool{
	var (
		suffix, domain, name, at bool
	)
	name = false
	domain = false
	suffix = false

	// Check wether there's a domain or not
	index_at_dot := strings.LastIndex(input, "@.")
	index_dot_co := strings.LastIndex(input, ".co.id")
	index_id := strings.LastIndex(input, ".id")
	if index_dot_co != -1 {
		if index_at_dot + 2 != index_dot_co {
			domain = true
		}
	} else if index_id != -1 {
		if index_at_dot + 2 != index_id {
			domain = true
		}
	}

	// Untuk Check panjang email maksimum 20 karakter sebelum, menggunakan index serta jumlah @ hanya boleh 1
	index_at := strings.LastIndexAny(input, "@")
	if index_at <= 20 && index_at == index_at_dot {
		name = true
	}

	// Check @.
	at = strings.Contains(input, "@.")

	// Check suffix
	if index_dot_co != -1 || index_id != -1 {
		suffix = true
	}

	return name && at && domain && suffix
}

func main(){
	var input string
	fmt.Print("Masukkan Input: ")
	fmt.Scanln(&input)
	fmt.Println("Email Valid: ", emailCheck(input))
}