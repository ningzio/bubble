package main

import (
	"fmt"

	"github.com/ningzio/bubble/internal/core"
)

func main() {
	// 音名 A
	a, _ := core.NewPitch(core.NoteA)
	// 音名 C
	c, _ := core.NewPitch(core.NoteC)

	fmt.Println(a.DistanceEasy(c))

	// C 升高八度
	c.High8()
	fmt.Println(a.DistanceEasy(c))

	// C 降低八度
	c.Low8()
	fmt.Println(a.DistanceEasy(c))

	// A 升高八度
	a.High8()
	fmt.Println(a.DistanceEasy(c))
}
