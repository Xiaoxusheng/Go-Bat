package test

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestClass(t *testing.T) {
	t1 := time.Date(time.Now().Year(), time.Now().Month()+2, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Now().Location()).Unix()
	t2 := time.Date(time.Now().Year(), 9, 4, 0, 0, 0, 0, time.Local).Unix()
	t3 := (t1 - t2) / (60 * 60 * 24 * 7)
	c := int(math.Ceil(float64((t1-t2)/(7*60*60*24)))) + 1
	w := time.Now().Weekday()

	fmt.Println(t1, t2, t3, c, w, int(w))
}
