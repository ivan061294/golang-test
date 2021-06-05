package controller

import (
	"encoding/json"
	"net/http"

	"github.com/JoseEvangelistaCucho/mod/entities"
)

func EntitiesArray(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var array entities.ArrayEntities
	if e := json.NewDecoder(r.Body).Decode(&array); e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	DatosArray := array.Rotate9()
	json.NewEncoder(w).Encode(DatosArray)
}
