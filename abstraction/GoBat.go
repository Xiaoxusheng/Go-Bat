package abstraction

// 抽象接口
type Bat interface {
	Controls(s any)
}

type GoBat struct {
	bat Bat
}

func (bat *GoBat) SetStrategy(B Bat) {
	bat.bat = B
}

func (bat *GoBat) Deal(s any) {
	bat.bat.Controls(s)
}
