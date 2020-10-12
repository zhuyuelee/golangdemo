package main

import "fmt"

/* interface 示例
类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

实现某个接口的类型（除了实现接口方法外）可以有其他的方法。

一个类型可以实现多个接口。

接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

*/
func main() {
	var simp simpler
	simp = new(simple)
	simp.Set(10)
	d := simp.Get()
	fmt.Print(d)
}

type simpler interface {
	Get() int
	Set(int)
}

type simple struct {
	data int
}

func (s *simple) Get() int {
	return s.data
}

func (s *simple) Set(d int) {
	s.data = d
}
