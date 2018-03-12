package main

import (
	"fmt"
	"tcverify"
)

func main() {
	resp, err := tcverify.Validate("xxxxxxxxxxx")
	fmt.Println(resp, err)

	tc := "xxxxxxxxxxx"
	name := "BARIŞ"
	surname := "ESEN"
	date := "1996"

	resp, err = tcverify.Check(tc, name, surname, date)
	fmt.Println(resp, err)
}
