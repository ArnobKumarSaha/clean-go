package abcd

import (
	"fmt"
	"strings"
)

func StringLearn() {
	s := "Hello世界"
	// len(s) is equivalent to len( []byte(s) )
	// %[n]v is called explicit argument indices
	// notation [n] immediately before the verb indicates that the nth one-indexed argument is to be formatted instead.
	fmt.Printf("%T %[1]v %d\n", s, len(s)) // len is 11 here, as chinese characters are taking 3 byte each

	// rune is synonym from int32 for characters
	// it is the logical representation
	r := []rune(s)
	fmt.Printf("%T %[1]v %d\n", r, len(r)) // len is 7 here, logical like this : 'H' 'e' 'l' 'l' 'o' '世' '界'
	// Output => []int32 [72 101 108 108 111 19990 30028] 7

	// rune is synonym from uint8
	// it is the physical representation
	b := []byte(s)
	fmt.Printf("%T %[1]v %d\n", b, len(b))
	// Output => []uint8 [72 101 108 108 111 228 184 150 231 149 140] 11

	// string is kind of Descriptor in go
	s = "hello, world"
	// Now, 's' knows two things only. 1) memory of the first byte,  2) length
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello, ",", world)
	// here `s`, `hello` & `world` are sharing the exact same underlying structure

	// as string is immutable, we can't do this => s[4] = 'w'
	// but, we can do this
	fmt.Printf("Before : %s %p\n", s, &s)
	s = s[:4] + "w" + s[5:]
	fmt.Printf("After : %s %p\n", s, &s)

	// it is illegal to take a particular element of a string
	// like this => fmt.Printf("%p\n", &s[0])

	// changing all the `world` occurrences to `universe`
	fmt.Println(strings.Join(strings.Split(s, "world"), "universe"))
}
