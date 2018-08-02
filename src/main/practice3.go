package main

import "fmt"

type I interface {
	I_Method()
}

type T struct {
	Name string
	Msg string
}

func (t T) SetMsg_ByVal(msg string) {
	//予備元には反映されない。副作用のない使い方
	t.Msg = msg
}
func (t *T) SetMsg_ByRef(msg string) {
	//予備元にも反映される副作用のある使い方
	t.Msg = msg
}
//I interface override Method
func (t *T) I_Method() {
	fmt.Printf("%v implements I concrete method\n", t.Name)
}

func main() {
	fmt.Println("START")

	//call type method
	b := T{"T1", "AAA"}
	b.SetMsg_ByVal("BBB")
	fmt.Println(b.Msg)
	b.SetMsg_ByRef("CCC")
	fmt.Println(b.Msg)

	a := T{"T2", ""}
	a.I_Method()

	//interface implement
	var i I
	i = &T{"T3", ""}
	i.I_Method()

	fmt.Println("END")
}
