package abcd

import "fmt"

func SliceLearn() {
	x := [...]int{1, 2, 3}
	y := []int{1, 2, 3}
	fmt.Printf("%T , %T\n", x, y) // [3]int , []int

	z := make([]int, 2)
	copy(z, y)
	// copy works on the minimum length, As len(z) is 2, len(y) is 3, it only copies 2 indices
	z[0] = 10 // z = [10, 2] now

	cx := x // copying an Array
	cx[0] = 5
	fmt.Println(x, cx) // [1 2 3] [5 2 3]

	cy := y // But this is not a copying, As `y` & `cy` is referring the same underlying structure. Use copy() instead
	cy[0] = 5
	fmt.Println(y, cy) // [5 2 3] [5 2 3]

	// slices are not comparable
	// this will give syntax error =>   if y == cy {}
	if x == cx {
		fmt.Println("Array comparing is Ok")
	}

	// assignment is not possible like,  `array = slice`  OR `slice = array`

	slice := []string{"a", "b", "c"}
	var ns []*string
	for idx, _ := range slice {
		ns = append(ns, &slice[idx])
	}

	for _, s := range ns {
		fmt.Println(*s)
	}
}
