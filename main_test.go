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
