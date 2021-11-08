package main

import (
	"fmt"
	"testProject/teal"
)
// a test
func main() {
	m := make(map[*interface{}]interface{})
	//var i int = 6
	var x, y interface{} = 3, 4
	var p *interface{}
	m[&x] = x
	m[&y] = y
	p = &x
	fmt.Printf("%v", x)
	*p = 10.7
	fmt.Printf("%v", x)

	var cb *teal.ConstByteArray = teal.NewConstByteArray([]byte{})
	res := cb.Equals(teal.NewConstByteArray([]byte{}))
	println(res)

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
