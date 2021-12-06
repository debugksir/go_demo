package example01

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello W3Cschool!")
	var strList []string
	fmt.Println(strList, len(strList), cap(strList), strList == nil) // 未初始化的变量等于nil

	var map01 = map[string]string{
		"a": "1",
		"b": "2",
	}
	fmt.Println(map01, map01["a"])
	fmt.Println(map01["c"] == "")
	d, ok := map01["d"] // ok可用于判断map是否存在key
	fmt.Println("d:", d, "--ok:", ok)

	a := [3]int{1, 2, 3}

	var map02 = make(map[string][]int)
	map02["a"] = a[:1]
	fmt.Println(map02["a"], map02["b"] == nil)

	type B struct {
		A string
		B int
		c bool
	}

	var b B
	b.A = "a"
	b.B = 1
	b.c = true

	fmt.Println(&b)

	var b1 = new(B)
	b1.A = "a"
	fmt.Println(b1)

	b2 := &B{"a", 1, true}
	fmt.Println(b2)

	var b3 *B = b1
	fmt.Println(b3)

	var b4 *B = &b
	fmt.Println(b4)
	fmt.Println(*b4)

	m := make(map[string]*student)
	stus := []student{
		{name: "ksir.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Println(m)

	n := make(map[string]student)
	for _, stu := range stus {
		n[stu.name] = stu
	}
	for k, v := range n {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Println(n)
}
