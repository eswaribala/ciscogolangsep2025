package main

import "github.com/brianvoe/gofakeit/v7"

func main() {

	min := 1000
	max := 9999
	otp := genOTP(&min, &max)
	println(otp)
	println(min)
	println(max)

}

func genOTP(min *int, max *int) int {
	*min = 15000
	*max = 19000
	return gofakeit.IntRange(*min, *max)
}
