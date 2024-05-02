package main

import "fmt"

type Animal struct {
	Nome       string
	Especie    string
	Felicidade int
}

type Recinto struct {
	EspeciesAnimais string
	Animais         []*Animal
	BemCuidado      bool
}

func NovoAnimal(nome, especie string, felicidade int) *Animal {
	return &Animal{
		Nome:       nome,
		Especie:    especie,
		Felicidade: felicidade,
	}
}

func NovoRecinto(especies string, bemCuidado bool) *Recinto {
	return &Recinto{
		EspeciesAnimais: especies,
		BemCuidado:      bemCuidado,
	}
}

func (r *Recinto) AdicionarAnimal(animal *Animal) {
	if animal.Especie == r.EspeciesAnimais {
		r.Animais = append(r.Animais, animal)
	}
}

func (a *Animal) Alimentar() {
	a.Felicidade += 10
}

func main() {
	fmt.Println("OOP em go chefe!")
}
