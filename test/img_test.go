package test

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_img(t *testing.T) {
	width := 1960
	height := 1080
	times := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(times)
	dc := gg.NewContext(width, height)
	dc.SetHexColor("#B3C890")
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	//字体
	face, err := gg.LoadFontFace("../config/t.ttf", 60)
	if err != nil {
		log.Panicln(err)
		return
	}
	f, err := gg.LoadFontFace("../config/t.ttf", 40)
	if err != nil {
		log.Panicln(err)
		return
	}
	dc.SetFontFace(face)
	dc.SetHexColor("#2D4356")
	//加载图片
	image, err := gg.LoadImage("../config/3.png")
	if err != nil {
		return
	}

	//时间
	dc.DrawStringAnchored(times, 210, float64(height-100), 0.5, 0.5)

	dc.DrawImageAnchored(image, width-170, height-80, 0.5, 0.5)

	dc.SetColor(color.RGBA{249, 251, 231, 150})
	dc.SetFontFace(f)

	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 10; i++ {
		fmt.Println()
		dc.Push()
		dc.RotateAbout(gg.Radians(40), float64(width/2), float64(height/2))
		dc.DrawStringAnchored("@GoBat", float64(rand.Int63n(1920)), float64(rand.Int63n(1080)), 0.5, 0.5)
		dc.Pop()
	}
	err = dc.SavePNG("../config/t.png")
	if err != nil {
		log.Panicln(err)
	}
}
