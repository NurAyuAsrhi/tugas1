package main

import (
	"fmt"
)

func main() {
	var a, z string
	var b, c, d, e float32
	a = "y"
	e = 1

	fmt.Println("Selamat menggunakan aplikasi operator aritmatika")

	for a != "n" {
		fmt.Print("Apakah anda akan melakukan perhitungan ( y / n) = ")
		fmt.Scan(&a)

		if a == "y" {
			fmt.Println("----------------------------------------------")
			fmt.Print("Input Operand 1 : ")
			fmt.Scan(&b)
			fmt.Print("Input Operator Aritmatika :")
			fmt.Scan(&z)
			fmt.Print("input Operand 2 : ")
			fmt.Scan(&c)

			if z == "+" {
				d = b + c
			} else if z == "-" {
				d = b - c
			} else if z == "*" {
				d = b * c
			} else if z == "/" {
				d = b / c
			} else if z == "^" {
				d = b
				for e < c {
					d = d * b
					e++
				}
			} else {
				fmt.Print("Operator tidak terdefinisi, mulai dari awal lagi ya.")
			}

			fmt.Print("Hasil")
			fmt.Print(b)
			fmt.Print(z)
			fmt.Print(c)
			fmt.Print(" = ")
			fmt.Println(d)
			fmt.Println("---------------------------------------------------")

		} else if a != "n" {
			fmt.Println("input hanya -y- untuk YES atau -N- untuk NO")

		}
	}
	fmt.Println("Terimakasih")
}
