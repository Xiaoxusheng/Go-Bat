package test

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func Test_m(t *testing.T) {
	var m = 8
	var n = &m
	*n = 100
	fmt.Println(&m, m, n, *n)
	var peo People = &Stduent{}
	think := "love"
	fmt.Println(peo.Speak(think))
}
