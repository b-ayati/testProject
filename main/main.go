package main

import (
	"testProject/teal"
)

func main() {
	var cb *teal.ConstByteArray = teal.NewConstByteArray([]byte{})
	res := cb.Equals(teal.NewConstByteArray([]byte{}))
	println(res)

	m := teal.NewMemorySegment(5)
	m.SaveSnapshot()
	m.DiscardSnapshot()
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
