package test

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func Test_img(t *testing.T) {
	str := ""
	var width float64
	var s string
	list := make([]string, 0)
	// 判断文件是否存在
	if _, err := os.Stat("../config/tf.png"); err != nil {
		//不存咋
		width := 1960
		height := 1080
		times := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(times)
		dc := gg.NewContext(width, height)
		dc.SetHexColor("#E1AEFF")
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
		dc.SetHexColor("#27374D")
		//加载图片
		image, err := gg.LoadImage("../config/3.png")
		if err != nil {
			return
		}

		//时间
		//dc.DrawStringAnchored(times, 210, float64(height-70), 0.5, 0.5)

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

	} else {
		//图片存在
		img, err := gg.LoadImage("../config/t.png")
		if err != nil {
			log.Panicln(err)
		}
		face, err := gg.LoadFontFace("../config/t.ttf", 70)
		if err != nil {
			log.Panicln(err)
			return
		}
		dc := gg.NewContextForImage(img)
		dc.SetFontFace(face)
		dc.SetHexColor("#333")
		wd, _ := dc.MeasureString(str)
		fmt.Println("wd", wd)
		if wd < 1960 {
			//不满1行
			dc.DrawString(str, 40, 350)
		} else {
			var h float64
			var w float64
			for i, r := range str {

				w, h = dc.MeasureString(string(r))
				wd, _ := dc.MeasureString(str[i:])
				width += w
				s += string(r)
				if width >= 1960 {
					list = append(list, s)
					width = 0
					s = ""
				}

				if wd < 1000 && i == len(str)-2 || i == len(str)-3 || i == len(str)-1 {
					list = append(list, s)
				}

			}
			for i, s := range list {
				fmt.Println(i, s)
				if i == 0 {
					dc.DrawString(s, 40, 80+float64(i)*h)
					continue
				}
				dc.DrawString(s, 0, 80+float64(i)*h)
			}
		}
		err = dc.SavePNG("../config/f.png")
		if err != nil {
			log.Panicln(err)
		}
	}
}
