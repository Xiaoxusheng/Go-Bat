package abstraction

// 抽象接口
type Bat interface {
	Controls(s any) any
}

type GoBat struct {
	bat Bat
}

// 设置策略
func (bat *GoBat) SetStrategy(B Bat) {
	bat.bat = B
}

// 调用
func (bat *GoBat) Deal(s any) any {
	return bat.bat.Controls(s)
}
