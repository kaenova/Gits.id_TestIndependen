/*
Gits.ID Assessment
Kaenova Mahendra Auditama
kaenova@gmail.com
*/

package main

import (
	"fmt"
	"strconv"
)

func konversiWaktu(jam int, menit int, detik int, waktu string) string{
	var (
		jam_conv int
		jam_final, menit_final string
	)


	if ( waktu == "PM" || waktu == "AM" ) && ( jam <= 12 && jam >= 0 ) && ( menit <= 59 && menit >= 0 )  && (detik <= 59 && detik >= 0 ){
		if waktu == "PM" {
			// Kalau PM tambahkan 12 kecuali untuk jam 12
			if jam != 12 {
				jam_conv = jam + 12
			} else {
				jam_conv = jam
			}
		} else if waktu == "AM" && jam == 12 {
			// Kalau AM yang berubah hanya ketika jam 12 malam, yaitu jam 00:00
			jam_conv = 0
		} else {
			// Kalau misalkan AM akan tetap
			jam_conv = jam
		}
	
		// Final string conversion
		if jam_conv < 10 {
			jam_final = "0" + strconv.Itoa(jam_conv)
		} else {
			jam_final = strconv.Itoa(jam_conv)
		}

		if menit < 10 {
			menit_final = "0" + strconv.Itoa(menit)
		} else {
			menit_final = strconv.Itoa(menit)
		}
	
		return jam_final+":"+menit_final
	
	}

	return "Format Waktu Tidak Valid"
}

func main(){
	var (
		jam, menit, detik int
		waktu string
	)

	fmt.Println("Masukkan Input: ")
	fmt.Scanf("%d:%d:%d %s", &jam, &menit, &detik, &waktu)
	
	fmt.Println("Hasil Konversi:",konversiWaktu(jam, menit, detik, waktu))

}