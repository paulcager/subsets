// +build no

package main

import (
	"fmt"
)

func main() {
	for n := 0; n < 15; n++ {
		fmt.Printf("%d\t", n)
		for i := 0; i < 1<<uint(n); i++ {
			ch := 'A'
			for bit := 1; bit < 0x100; bit <<= 1 {
				if i&bit != 0 {
					fmt.Printf("%c", ch)
				}
				ch++
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
