package main

// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//     return f();
// }
//
// int fortytwo()
// {
//     return 42;
// }
import "C"
import "fmt"

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(C.bridge_int_func(f))

	p := &[3]int{1, 2, 3}
	pow(p)
	fmt.Println(p)

	p1 := &[3]string{"Apple", "Banana", "Cherry"}
	for i, v := range p1 {
		fmt.Println(i, v)
	}

	i := 5
	ip := &i
	fmt.Printf("type=%T, address=%p\n", ip, ip)

	s := ""
	for _, v := range []string{"A", "B", "C"} {
		s += v
	}
	fmt.Println(s)
}
