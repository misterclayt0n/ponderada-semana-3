package main

import (
	"reflect"
	"testing"
)

func TestNovoAnimal(t *testing.T) {
	animal := NovoAnimal("Simba", "Leão", 50)
	expected := &Animal{"Simba", "Leão", 50}
	if !reflect.DeepEqual(animal, expected) {
		t.Errorf("esperado %v, recebido %v", expected, animal)
	}
}

func TestAdicionarAnimal(t *testing.T) {
	recinto := NovoRecinto("Leão", true)
	animal := NovoAnimal("Simba", "Leão", 50)
	recinto.AdicionarAnimal(animal)
	if len(recinto.Animais) != 1 || recinto.Animais[0] != animal {
		t.Errorf("falha ao adicionar animal na gaiola")
	}
}

func TestAlimentarAnimal(t *testing.T) {
	animal := NovoAnimal("Simba", "Leão", 50)
	animal.Alimentar()
	if animal.Felicidade != 60 {
		t.Errorf("o nivel esperado de felicidade era 60, obtido: %d", animal.Felicidade)
	}
}

func TestAtrairVisitantesBemCuidado(t *testing.T) {
	recinto := NovoRecinto("Leão", true)
	animal := NovoAnimal("Simba", "Leão", 60)
	recinto.AdicionarAnimal(animal)
	recinto.AtrairVisitantes()
	if recinto.Visitantes != 10 {
		t.Errorf("esperado 10 visitantes, recebido %d", recinto.Visitantes)
	}
}

func TestAtrairVisitantesMalCuidado(t *testing.T) {
	recinto := NovoRecinto("Leão", false)
	animal := NovoAnimal("Simba", "Leão", 60)
	recinto.AdicionarAnimal(animal)
	recinto.AtrairVisitantes()
	if recinto.Visitantes != 0 {
		t.Errorf("esperado 0 visitantes, recebido %d", recinto.Visitantes)
	}
}

func TestCalcularFelicidadeMediaVazia(t *testing.T) {
	recinto := NovoRecinto("Leão", true)
	if media := recinto.calcularFelicidadeMedia(); media != 0 {
		t.Errorf("esperado 0 de felicidade média, recebido %d", media)
	}
}

func TestCalcularFelicidadeMediaComAnimais(t *testing.T) {
	recinto := NovoRecinto("Leão", true)
	animal1 := NovoAnimal("Simba", "Leão", 40)
	animal2 := NovoAnimal("Nala", "Leão", 60)
	recinto.AdicionarAnimal(animal1)
	recinto.AdicionarAnimal(animal2)
	expectedMedia := 50
	if media := recinto.calcularFelicidadeMedia(); media != expectedMedia {
		t.Errorf("esperado %d de felicidade média, recebido %d", expectedMedia, media)
	}
}
