package main

import (
	"awesomeProject/decorator/pkg"
	. "fmt"
)

func main() {
	base := pkg.BasePc{}

	homePc := pkg.HomePc{
		Cpu:     4,
		Gpu:     1,
		Wrapper: base,
	}

	serverPc := pkg.ServerPC{
		Cpu:     8,
		Memory:  32,
		Wrapper: base,
	}

	Println(base)
	Println(homePc)
	Println(serverPc)

	Println(base.GetPrice())
	Println(homePc.GetPrice())
	Println(serverPc.GetPrice())
}
