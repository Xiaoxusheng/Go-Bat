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
	"unicode"
)

type Picture struct {
}

func (p *Picture) CreatePicture(strs string) {
	t := time.Now()
	str := strings.ReplaceAll(strs, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "   ", "")
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
		//内容字体
		face, err := gg.LoadFontFace("./config/lishu.ttf", 50)
		if err != nil {
			log.Panicln(err)
		}
		//水印字体
		f, err := gg.LoadFontFace("./config/t.ttf", 70)
		if err != nil {
			log.Panicln(err)
			return
		}
		dc := gg.NewContextForImage(img)
		dc.SetColor(color.RGBA{249, 251, 231, 150})
		dc.SetFontFace(f)

		rand.Seed(time.Now().UnixMicro())
		for i := 0; i < 10; i++ {
			fmt.Println()
			dc.Push()
			dc.RotateAbout(gg.Radians(45), float64(1096/2), float64(1080/2))
			dc.DrawStringAnchored("@GoBat", float64(rand.Int63n(1920)), float64(rand.Int63n(1080)), 0.5, 0.5)
			dc.Pop()
		}
		//内容
		dc.SetFontFace(face)
		dc.SetHexColor("#E21818")
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

				// (i == len(str)-3 && !unicode.IsLetter(rune(str[i : i+1][0])))意思为当i-3时如果不是中文继续写入，是中文才添加到列表中
				//(i == len(str)-3 && !unicode.IsSymbol(rune(str[i : i+1][0])))意思为当i-3时如果不是字符=这种继续写入，是中文才添加到列表中
				if wd < 1800 && (i == len(str)-3 && !unicode.IsLetter(rune(str[i : i+1][0])) && !unicode.IsSymbol(rune(str[i : i+1][0])) && !unicode.IsNumber(rune(str[i : i+1][0])) && !unicode.IsSpace(rune(str[i : i+1][0]))) || i == len(str)-1 {
					fmt.Println(i, len(str), s, unicode.IsLetter(rune(str[i : i+1][0])), unicode.IsSymbol(rune(str[i : i+1][0])), unicode.IsSpace(rune(str[i : i+1][0])), str[i:i+1])
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
		//时间字体
		dc.SetFontFace(h)
		dc.DrawStringAnchored(times, 210, float64(1080-70), 0.5, 0.5)
		err = dc.SavePNG("./config/f.png")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println("生成完成,时间为：", time.Now().Sub(t))
	}
}
