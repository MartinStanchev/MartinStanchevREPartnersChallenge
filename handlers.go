package main

import (
	"encoding/json"
	"net/http"
)

type order struct {
	Order int `json:"order"`
}

type packSizes struct {
	PackSizes []int `json:"packSizes"`
}

func (app *application) handleOrder(w http.ResponseWriter, r *http.Request) {
	var reqOrder order
	err := json.NewDecoder(r.Body).Decode(&reqOrder)
	if err != nil {
		app.log.WithError(err).Warn("could not unmarshal body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.log.WithField("order", reqOrder.Order).Info("new order")

	returnPacks, err := app.distributor.Distribute(reqOrder.Order)
	if err != nil {
		app.log.WithError(err).Errorf("could not calculate return packs")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(returnPacks)
	if err != nil {
		app.log.WithError(err).Errorf("could not marshal response to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) updatePackSizes(w http.ResponseWriter, r *http.Request) {
	var packSizes packSizes
	err := json.NewDecoder(r.Body).Decode(&packSizes)
	if err != nil {
		app.log.WithError(err).Warn("could not unmarshal body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(packSizes.PackSizes) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.distributor.PackSizes = packSizes.PackSizes

	app.log.WithField("PackSizes", packSizes.PackSizes).Info("updating pack sizes")

	w.WriteHeader(http.StatusOK)
}
