package main

import (
	"testing"
)

func TestMultiplicacao(t *testing.T) {
	casesM := []struct {
		Name string
		A    int
		B    int
		Exp  int
	}{
		{
			Name: "TEST 2 * 8",
			A:    2,
			B:    8,
			Exp:  16,
		},
		{
			Name: "TEST 3 * 8",
			A:    3,
			B:    8,
			Exp:  24,
		},
	}

	for _, m := range casesM {
		t.Run(m.Name, func(t *testing.T) {
			testCalculo(t, m.A, m.B, m.Exp)
		})
	}
}

func testCalculo(t *testing.T, a, b, exp int) {
	got := multiplica(a, b)
	if got != exp {
		t.Errorf("Houve um erro de calculo. esperado [%d], recebido [%d]", exp, got)
	}
}

func TestVariadicos(t *testing.T) {
	casesM := []struct {
		Name string
		C    []int
		Exp  int
	}{

		{
			Name: "TEST Slice int",
			C:    []int{1, 2, 3, 4, 5},
			Exp:  15,
		},
	}

	for _, n := range casesM {
		t.Run(n.Name, func(t *testing.T) {
			testvariadico(t, n.Exp, n.C...)
		})
	}
}

func testvariadico(t *testing.T, exp int, i ...int) {
	got := exemploVariadico(i...)
	if got != exp {
		t.Errorf("Houve um erro de calculo. esperado [%d], recebido [%d]", exp, got)
	}
}

func Test_exemploNomeado(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name     string
		args     args
		wantNome string
	}{
		{//caso 1, esse deve ser IGUAL a estrutura acima
			"TestOk",
			args{"Teste"},
			"Teste",
		},
		{//caso 2, esse pode inverter a ordem dos tipos
			name:"TestNOk",
			args:args{"igual"},
			wantNome:"igual",
		},
		{//caso 2, esse pode inverter a ordem dos tipos mudados
			wantNome:"igual",
			args:args{"igual"},
			name:"TestNOk",
			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNome := exemploNomeado(tt.args.str); gotNome != tt.wantNome {
				t.Errorf("exemploNomeado() = %v, want %v", gotNome, tt.wantNome)
			}
		})
	}
}
