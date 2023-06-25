package config

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Picture struct {
}

func (p *Picture) CreatePicture(strs string) {
	str := strings.ReplaceAll(strs, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, " ", "")
	var width float64
	var s string
	list := make([]string, 0)
	times := time.Now().Format("2006-01-02 15:04:05")
	// 判断文件是否存在
	if _, err := os.Stat("./config/t.png"); err != nil {
		//不存在
		width := 1960
		height := 1080
		fmt.Println(times)
		dc := gg.NewContext(width, height)
		dc.SetHexColor("#B3C890")
		dc.DrawRectangle(0, 0, float64(width), float64(height))
		dc.Fill()
		//字体
		face, err := gg.LoadFontFace("./config/t.ttf", 60)
		if err != nil {
			log.Panicln(err)
			return
		}
		f, err := gg.LoadFontFace("./config/t.ttf", 40)
		if err != nil {
			log.Panicln(err)
			return
		}
		dc.SetFontFace(face)
		dc.SetHexColor("#30A2FF")
		//加载图片
		image, err := gg.LoadImage("./config/3.png")
		if err != nil {
			return
		}

		//时间
		dc.DrawStringAnchored(times, 210, float64(height-70), 0.5, 0.5)

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
		err = dc.SavePNG("./config/t.png")
		if err != nil {
			log.Panicln(err)
		}

	} else {
		//图片存在

		img, err := gg.LoadImage("./config/t.png")
		if err != nil {
			log.Panicln(err)
		}
		face, err := gg.LoadFontFace("./config/4.ttf", 50)
		if err != nil {
			log.Panicln(err)
		}
		dc := gg.NewContextForImage(img)
		dc.SetFontFace(face)
		dc.SetHexColor("#EEE3CB")
		wd, _ := dc.MeasureString(str)
		if wd < 1900 {
			//不满1行
			dc.DrawString(str, 40, 80)
		} else {
			var h float64
			var w float64
			for i, r := range str {
				w, h = dc.MeasureString(string(r))
				wd, _ := dc.MeasureString(str[i:])
				width += w
				s += string(r)
				if width >= 1900 {
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
		h, err := gg.LoadFontFace("./config/t.ttf", 60)
		if err != nil {
			log.Panicln(err)
			return
		}
		dc.SetHexColor("#27374D")

		dc.SetFontFace(h)
		dc.DrawStringAnchored(times, 210, float64(1080-70), 0.5, 0.5)
		err = dc.SavePNG("./config/f.png")
		if err != nil {
			log.Panicln(err)
		}
	}
}
