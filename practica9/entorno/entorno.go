package entorno

import (
	"fmt"
	"strconv"
)

type Entorno struct {
	Ant   *Entorno
	Tabla map[string]*Simbolo
}

type IEntorno interface {
	Put(string, *Simbolo)
	Get(string)
}

func NewEntorno(ant *Entorno) *Entorno {
	return &Entorno{Ant: ant, Tabla: make(map[string]*Simbolo)}
}

func (ent *Entorno) Put(key string, sim *Simbolo) bool {
	_, encontrado := ent.Tabla[key]

	if !encontrado {
		ent.Tabla[key] = sim
		return true
	}
	return false
}

func (ent *Entorno) Get(key string) *Simbolo {
	for e := ent; e != nil; e = e.Ant {
		sim, ok := e.Tabla[key]

		if ok {
			return sim
		}
	}
	return nil
}

// mostrar símbolos de un entorno x
func Mostrar(ent *Entorno) {
	for _, v := range ent.Tabla {
		fmt.Print("Identificador: ", v.Id)
		fmt.Print("\tTipo: ", v.Tipo)
		fmt.Println("\tDireccion: ", v.Dir)
	}
}

// manejo de ambitos como lo maneja el libro
var pila []*Entorno

func Push(e *Entorno) {
	pila = append(pila, e)
}

func Pop() *Entorno {
	if len(pila) < 1 {
		panic("empty env")
	}
	result := pila[len(pila)-1]
	pila = pila[:len(pila)-1]
	return result
}

func GetTamDim(id string, noDim int, e *Entorno) string {
	tam := 1

	s := e.Get(id)

	for i := 0; i < noDim; i++ {
		dim := s.Dims[i]
		dif := (dim.Lsup - dim.Linf) + 1
		tam = tam * dif
	}

	return strconv.Itoa(tam)
}

func GetTamDimFilas(id string, noDim int, e *Entorno) string {
	tam := 1

	s := e.Get(id)
// [3][10][5]
	for i := noDim; i <  len(s.Dims); i++ {
		dim := s.Dims[i]
		dif := (dim.Lsup - dim.Linf) + 1
		tam = tam * dif
	}
	return strconv.Itoa(tam)
}

