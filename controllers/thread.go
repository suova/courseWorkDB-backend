package controllers

import (
	"databasework/models"
	"databasework/queries"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandleThreadGet(w http.ResponseWriter, r *http.Request) {

	result := queries.QueriesGetAllThreads()
	resp, _ := json.Marshal(result)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleCreateThread(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	thread := &models.Thread {}
	if err := json.Unmarshal(body, thread); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queries.CreateThread(thread)
	j, _ := json.Marshal(thread)
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

