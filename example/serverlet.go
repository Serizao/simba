package main

import (
	"fmt"

	"github.com/Serizao/simba"
)

func main() {
	fmt.Printf("Hello, world.")

	// Listen 445 Port

	s := &simba.Server{}
	s.ListenAndServe("0.0.0.0:1445")

}
