package main

import "testing"

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
	t.Fatal("not implemented")
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
