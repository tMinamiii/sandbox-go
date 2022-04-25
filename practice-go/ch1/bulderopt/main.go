package main

import "fmt"

type Portion int

const (
	Regular Portion = iota // 普通
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion) *fluentOpt {
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) String() {
	fmt.Printf("面: %d, 油揚げ: %v, 海老天: %d\n", o.men, o.aburaage, o.ebiten)
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func useFluentInterface() {
	oomoriKitsune := NewUdon(Large).Aburaage().Order()
	fmt.Println(oomoriKitsune)
}
