package main

import (
	"fmt"
	"testProject/teal"
)

func main() {
	var arr [3]teal.DataType
	arr[0] = teal.NewUInt(5)
	fmt.Printf("%v", arr)
	if arr[1] == nil {
		println("hi")
	}

	/*
		x := teal.NewUInt(25)
		y := teal.NewUInt(23)

		m := new(teal.MemorySegment)
		m.SaveSnapshot()
		m.Add(x)
		m.Add(y)
		x.SetValue(19)
		m.RestoreSnapshot()
		fmt.Print(m)
		//fmt.Printf("(%v, %T)\n", inter1, inter1)
	*/
}
