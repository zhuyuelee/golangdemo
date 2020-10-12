package main

import "fmt"

//Shaper 接口
type Shaper interface {
	Area() float32
}

//Square 正方形
type Square struct {
	side float32
}

//Area 面积
func (sq Square) Area() float32 {
	return sq.side * sq.side
}

//Rectangle 长方形
type Rectangle struct {
	length, width float32
}

//Area 面积
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	m := make(map[string]Shaper)
	m["Rectangle"] = Rectangle{5, 3}
	m["Square"] = Square{10}

	for key, s := range m {
		fmt.Printf("%s的面积是%f\n", key, s.Area())

	}

	// r := Rectangle{5, 3} // Area() of Rectangle needs a value
	// q := &Square{5}      // Area() of Square needs a pointer
	// // shapes := []Shaper{Shaper(r), Shaper(q)}
	// // or shorter
	// shapes := []Shaper{r, q}
	// fmt.Println("Looping through shapes for area ...")
	// for n, _ := range shapes {
	// 	fmt.Println("Shape details: ", shapes[n])
	// 	fmt.Println("Area of this shape is: ", shapes[n].Area())
	//}
}
