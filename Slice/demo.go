package main

import (
	"fmt"
)

func main() {
	//Slice相关细节，首先如果要定义一个切片，需要使用make函数
	//func make(t Type, size ...IntegerType) Type
	s1 := make([]int, 5, 10) //长度为5，容量为10
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

}
