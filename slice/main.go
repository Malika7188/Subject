package main

import "fmt"

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
func Slice(a []string, nbrs ...int) []string {
	if len(nbrs)==1{
		if nbrs[0] < 0 {
			a=rev(a)
			nbrs[0]= -nbrs[0]
			return (a[nbrs[0]:])
		}
		return a[nbrs[0]:]
	}

	if nbrs[0] < 0 && nbrs[1] < 0 {
		nbrs[0] = -nbrs[0]
		nbrs[1] = -nbrs[1]
	}
	if nbrs[0] < 0 {
		nbrs[0] = -nbrs[0]
	}
	if nbrs[1] < 0 {
		nbrs[1] = -nbrs[1]
	}

	first := nbrs[0]
	
	sec := nbrs[1]
	return a[first:sec]
}

func rev(a []string) []string{
	res := []string{}
	for i := len(a)-1; i >= 0; i--{
		res=append(res,a[i])
	}
	return res
}