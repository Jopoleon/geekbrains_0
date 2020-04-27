package controllers

import (
	"encoding/json"
	"net/http"
)

func (a *Controllers) GetAuthors(w http.ResponseWriter, r *http.Request) {
	res, err := a.Repository.DB.GetAllAuthors()
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

func (a *Controllers) CreateAuthor(w http.ResponseWriter, r *http.Request) {

}
