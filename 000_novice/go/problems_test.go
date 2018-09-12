package main

import (
	"reflect"
	"testing"
)

func TestNegue(t *testing.T) {
	var tests = []struct {
		input bool
		want  bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range tests {
		if got := Negue(test.input); got != test.want {
			t.Errorf("Negue(%t) = %t", test.input, got)
		}
	}
}

func TestDigaOla(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", "Olá!"},
		{"    ", "Olá!"},
		{"Paulo", "Olá, Paulo."},
		{"  Paulo  ", "Olá, Paulo."},
	}

	for _, test := range tests {
		if got := DigaOla(test.input); got != test.want {
			t.Errorf("DigaOla(%q) = %q", test.input, got)
		}
	}
}

func TestListaNumerosPares(t *testing.T) {
	var tests = []struct {
		input int
		want  []int
	}{
		{0, []int{}},
		{1, []int{2}},
		{4, []int{2, 4, 6, 8}},
		{5, []int{2, 4, 6, 8, 10}},
		{-1, []int{}},
	}

	for _, test := range tests {
		got := ListaNumerosPares(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ListaNumerosPares(%v) = %v", test.input, got)
		}
	}
}

func TestListaMultiplos(t *testing.T) {
	t.Fatal("not implemented")
}

func TestSoma(t *testing.T) {
	t.Fatal("not implemented")
}

func TestSubtracao(t *testing.T) {
	t.Fatal("not implemented")
}

func TestMultiplicacao(t *testing.T) {
	t.Fatal("not implemented")
}

func TestDivisao(t *testing.T) {
	t.Fatal("not implemented")
}

func TestOperacao(t *testing.T) {
	t.Fatal("not implemented")
}

func TestMaior(t *testing.T) {
	t.Fatal("not implemented")
}

func TestIntersecao(t *testing.T) {
	t.Fatal("not implemented")
}
