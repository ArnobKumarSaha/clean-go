package go_basics

import "fmt"

func FormatLearn() {
	a, b := 12, 345
	c, d := 1.2, 3.45

	fmt.Printf("|%06d|%6d| \n", a, b)     // |000012|   345|
	fmt.Printf("|%-6d|%-06d| \n", a, b)   // |12    |345   |
	fmt.Printf("%d  %x  %X  \n", a, a, a) // 12  c  C
	fmt.Printf("%#x  %#x \n", a, b)       // 0xc  0x159
	fmt.Printf("%.2f  %.4f \n", c, d)     // 1.20  3.4500
	fmt.Printf("%6f  %6.2f \n", c, d)     // 1.200000    3.45

	s := []int{1, 2}
	fmt.Printf("%T , %v , %#v \n", s, s, s) // []int , [1 2] , []int{1, 2}

	r := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%T , %q , %v ,%#v \n", r, r, r, r) // [3]int32 , ['a' 'b' 'c'] , [97 98 99] ,[3]int32{97, 98, 99}

	m := map[string]int{"x": 1, "y": 2}
	fmt.Printf("%T , %v , %#v \n", m, m, m) // map[string]int , map[x:1 y:2] , map[string]int{"x":1, "y":2}

	t := "hi"
	bs := []byte(t)
	fmt.Printf("%T , %v , %#v , %q \n", t, t, t, t)     // string , hi , "hi" , "hi"
	fmt.Printf("%T , %v , %#v , %v \n", bs, bs, bs, bs) // []uint8 , [104 105] , []byte{0x68, 0x69} , [104 105]
}
