package main

import "fmt"

func main() {
	var list = make(map[string]Shaper)

	list["Square"] = &Square{Side: 10}
	list["Rectangle"] = &Rectangle{Width: 3, Height: 5}
	list["Circle"] = &Circle{Radius: 5}

	for key, s := range list {
		fmt.Printf("%s的页面是%0.2f\n", key, s.Area())
	}
}

//Shaper 接口
type Shaper interface {
	// Area 计算面积
	Area() float32
}

//Square 正方形
type Square struct {
	Side float32
}

// Area 计算面积
//指针类型的参数 必须传指针类型的实参
//如 $struce
func (s *Square) Area() float32 {
	return s.Side * s.Side
}

//Rectangle 长方形
type Rectangle struct {
	// Width 宽
	Width float32
	// Height 高
	Height float32
}

// Area 计算面积
func (r *Rectangle) Area() float32 {
	return r.Width * r.Height

}

//Circle 圆
type Circle struct {
	Radius float32
}

// Area 计算面积
func (c *Circle) Area() float32 {
	return c.Radius * c.Radius * 3.14
}
