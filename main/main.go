package main

import (
	"fmt"
	"testProject/teal"
)

func main() {

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

}
