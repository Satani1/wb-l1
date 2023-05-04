package main

import "fmt"

// целевой интерфейс, под который необходимо адаптироваться
type Person interface {
	SpeakEnglish()
}

type EnglishMan struct{}

func (em *EnglishMan) SpeakEnglish() {
	fmt.Println("Hello. *British speaks*")
}

// структура, которую необходимо адаптировать под интерфейс Person
type RussianMan struct{}

func (rm *RussianMan) SpeakRussian() {
	fmt.Println("Привет. *Говорит по-русски*")
}

// адаптер
type RussianManAdapter struct {
	Russian *RussianMan
}

func (rma *RussianManAdapter) SpeakEnglish() {
	fmt.Println("Hello. *Russian speaks on english*")
}

func main() {
	var englishMan EnglishMan
	var adapterRussianMan RussianManAdapter
	var russianMan RussianMan

	englishMan.SpeakEnglish()
	russianMan.SpeakRussian()
	adapterRussianMan.Russian.SpeakRussian() //адаптер может говорить на русском.
	adapterRussianMan.SpeakEnglish()         //так и на английском, поэтому и соответствует интерфейсу.
}
