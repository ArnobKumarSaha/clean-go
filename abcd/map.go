package abcd

import "fmt"

func MapLearn() {
	var m map[string]int // it refers to nothing (no hashtable referred)
	fmt.Println(m["hi"])
	// reading is OK, but below assignment will panic
	// m["hi"] = 2
	// use below syntax to avoid that panic
	m = make(map[string]int) // but, it refers to a hash table, which is empty
	m["hi"] = 2

	// assigning of map doesn't make a copy of itself
	// So, below `p` & `m` is exactly same, updating into one changes the other
	fmt.Println(m)
	p := m
	p["hi"]++
	fmt.Println(m)

	// map can't be compared to one another, it only can be compared to nil
	x := make(map[int]int)
	if x == nil {
	} // this condition will not be true
	var y map[int]int
	if y == nil {
	} // this condition will be true
	// this will be syntax error =>   if x == y {}

	// accessing a map gives two variable like this, VALUE, IsThere := mp[KEY]
}
