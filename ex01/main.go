package main

import "fmt"

// Human структура
type Human struct {
	Name string
}

// Say метод структуры Human
func (h Human) Say() {
	fmt.Println("Hello, i am", h.Name)
}

// Action структура
// Если у структуры Action нет метода Say, то она реализует метод Say структуры Human
type Action struct {
	Human
}

func main() {
	a := Action{
		Human{
			Name: "Tom",
		}}
	/*
		Метод Say структуры Human
		К методу Say структуры Human обращаемся через структуру Action,
		так же можно обращаться через структуру Human (a.Human.Say())
	*/
	a.Say()
}
