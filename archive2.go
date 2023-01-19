package main

import (
	"fmt"
	"sort"
	"unsafe"
)

type A struct {
	B int
}

func (a *A) megatest() {
	fmt.Println("This is A")
}

func unsafePointersTask() {
	foo := Foo{Bar: "123", Qux: 7, Dem: true}
	fooBytes := *(*[FooSize]byte)(unsafe.Pointer(&foo))
	barBytes := *(*[]byte)(unsafe.Pointer(&foo.Bar))
	pnt := &foo
	fmt.Println(fooBytes, barBytes)
	fmt.Printf("%p %T", pnt, fooBytes)
}

type Foo struct {
	Bar string
	Qux int
	Dem bool
}

type Foos []Foo

func (f Foos) Len() int           { return len(f) }
func (f Foos) Less(i, j int) bool { return f[i].Bar < f[j].Bar }
func (f Foos) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func sortFoosByBars() {
	foos := Foos{
		{Bar: "z"},
		{Bar: "d"},
		{Bar: "q"},
		{Bar: "a"},
		{Bar: "x"},
		{Bar: "r"},
		{Bar: "c"},
	}
	fmt.Println(foos)

	sort.Sort(foos)
	fmt.Println(foos)

}

const FooSize = unsafe.Sizeof(Foo{})

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)

	z := append(x, 4)
	xy, xz := &y, &z
	fmt.Println(y, z)
	fmt.Print(xy, xz)
}
