package main

import (
	"fmt"
)

func main() {
	var a string = "字符串"
	var b int = 1
	var c float64 = .2233
	var (
		name string = "Jesxion"
		age int = 32
	)
	d := "短类型声明"
	e, f := "声明和初始化多个变量", 1.2
	c, f = f, c
	var ptr = &name
	ptr2 := new(int)
	x, _ := 100, 200
	_, y := 100, 200
	fmt.Println(a, b, c, name, ptr, *ptr, *ptr2, age, d, e, f, x, y)
}