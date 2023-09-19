package test

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestClass(t *testing.T) {
	t1 := time.Now().Unix()
	t2 := time.Date(time.Now().Year(), 9, 4, 0, 0, 0, 0, time.Local).Unix()
	t3 := (t1 - t2) / (60 * 60 * 24 * 7)
	c := int64(math.Ceil(float64((t1-t2)/(7*60*60*24)))) + 1
	w := time.Now().Weekday()

	fmt.Println(t1, t2, t3, c, w, int(w))
}
