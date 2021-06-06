package entities

import (
	"github.com/ivan061294/golang-test/models"
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
		if nmatriz > n {
			r.Code = 1
			r.Message = "El numero de campos no es igual a los campos array"
			return r
		}
		if nmatriz < n {
			r.Code = 1
			r.Message = "El numero de array no es igual a los campos "
			return r
		}

		r.Code = 1
		r.Message = "Ha Ocurrido Un Error interno"
		return r
	}
	if nmatriz < n {
		r.Code = 1
		r.Message = "El numero de campos no es igual a los campos array ejm NxN"
		return r
	}
	if nmatriz > n {
		r.Code = 1
		r.Message = "El numero de array no es igual a los campos ejm NxN"
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
	r.Message = "Procesado Con Éxito"
	r.Data = ar

	return r
}
func (response Response) ResponseAll() (r Response) {

	response.Code = 1
	response.Message = "Este no es un tipo de Petición Valida. " +
		"La Petición es de tipo POST y la uri es"
	return response
}
