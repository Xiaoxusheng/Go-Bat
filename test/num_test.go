package test

import (
	"fmt"
	"testing"
	"unicode"
)

func Test_num(t *testing.T) {
	str := "idji id1"
	for _, i2 := range str {
		if unicode.IsNumber(i2) {
			fmt.Println(string(i2))
		}
	}

}
