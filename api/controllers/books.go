package controllers

import (
	"encoding/json"
	"net/http"
)

func (a *Controllers) GetBooks(w http.ResponseWriter, r *http.Request) {
	res, err := a.Repository.DB.GetAllBooks()
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) CreateBook(w http.ResponseWriter, r *http.Request) {

}
