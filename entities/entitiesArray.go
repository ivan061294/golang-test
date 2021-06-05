package entities

import (
	"github.com/JoseEvangelistaCucho/mod/models"
)

type ArrayEntities models.ArrayNum

type ArrayEntitiesrotate models.ArrayNumRotate
type Response models.Response

func (c *ArrayEntities) Rotate9() (r Response) {
	var a [][]int
	var m = c.ValorArray
	nmatriz := len(m)
	n := len(m[0])
	if nmatriz != n {
		r.Code = 1
		r.Message = "El numero de campos no es igual a los campos array"
		return r
	}
	a = make([][]int, n) // Crea la nueva "matriz"
	for i := range a {
		a[i] = make([]int, n) // y sus filas
	}

	//Intercambia las filas a columnas (intercambiando los índices j -> i, i -> j desde el final)
	for i := range m {
		for j := range m[i] {
			a[j][i] = m[i][n-j-1]
		}
	}

	var ar models.ArrayNumRotate
	ar.ValorArrayRotate = append(ar.ValorArrayRotate, a...)
	r.Code = 0
	r.Message = "Procesado Con Exito"
	r.Data = ar

	return r
}
