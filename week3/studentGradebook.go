package main

import "fmt"

func main() {

	gradebook := make(map[string]float64)

	for {
		fmt.Println("Enter pair (std spc grd):")
		var tmp_std string
		fmt.Scan(&tmp_std)
		var tmp_grd float64
		fmt.Scan(&tmp_grd)
		gradebook[tmp_std] = tmp_grd
		fmt.Println(gradebook)
	}
}
