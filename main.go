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
	Visitantes      int
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

func (a *Animal) String() string {
	return fmt.Sprintf("Nome: %s, Especie: %s, Felicidade: %d", a.Nome, a.Especie, a.Felicidade)
}

func (r *Recinto) String() string {
	return fmt.Sprintf("EspeciesAnimais: %s, Animais: %v, BemCuidado: %t, Visitantes: %d", r.EspeciesAnimais, r.Animais, r.BemCuidado, r.Visitantes)
}

func (r *Recinto) AdicionarAnimal(animal *Animal) {
	if animal.Especie == r.EspeciesAnimais {
		r.Animais = append(r.Animais, animal)
	}
}

func (a *Animal) Alimentar() {
	a.Felicidade += 10
}

func (r *Recinto) AtrairVisitantes() {
	if r.BemCuidado && r.calcularFelicidadeMedia() >= 50 {
		r.Visitantes = 10
	} else {
		r.Visitantes = 0
	}
}

func (r *Recinto) calcularFelicidadeMedia() int {
	if len(r.Animais) == 0 {
		return 0
	}
	totalFelicidade := 0
	for _, animal := range r.Animais {
		totalFelicidade += animal.Felicidade
	}
	return totalFelicidade / len(r.Animais)
}

func main() {
	fmt.Println("OOP em Go!")
}
