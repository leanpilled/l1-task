package main

import "fmt"

type Lada struct{}

func (l *Lada) Fill(gas string) {
	fmt.Printf("filled with %s\n", gas)
}

type Tesla struct{}

func (t *Tesla) Charge() {
	fmt.Println("charged battery")
}

type CarAdapter interface {
	Charge()
}

type LadaAdapter struct {
	*Lada
}

func (adapter *LadaAdapter) Charge() {
	adapter.Fill("92")
}

func NewLadaAdapter(car *Lada) *LadaAdapter {
	return &LadaAdapter{car}
}

func main() {
	lada := NewLadaAdapter(&Lada{})
	lada.Charge()
	tesla := &Tesla{}
	tesla.Charge()
}
