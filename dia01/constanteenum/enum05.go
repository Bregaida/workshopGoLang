package main

import (
	"fmt"
)

const (
	_                   = iota
	estrelaHiperGigante = 1 << iota
	estrelaSuperGigante
	estrelaBrilhanteGigante
	estrelaGigante
	estrelaSubGigante
	estrelaSubGigantesca
)

func main() {
	fmt.Println(estrelaHiperGigante)
	fmt.Println(estrelaSuperGigante)
	fmt.Println(estrelaBrilhanteGigante)
	fmt.Println(estrelaGigante)
	fmt.Println(estrelaSubGigante)
	fmt.Println(estrelaSubGigantesca)
}