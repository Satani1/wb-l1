package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetAge(age int) {
	h.Age = age
}

type Action struct {
	Human
}

func (a *Action) SetName(name string) {
	a.Name = name
}
func (a *Action) SetAge(age int) {
	a.Age = age
}

func main() {
	a1 := Action{
		Human{
			Name: "Test",
			Age:  20,
		},
	}

	fmt.Printf("Action struct at start: %+v\n", a1)

	a1.SetName("New Name")
	a1.SetAge(50)

	fmt.Printf("Action struct after methods SetName and SetAge: %+v\n", a1)
}
