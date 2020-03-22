package controllers

import (
	"databasework/models"
	"databasework/queries"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func HandleCreateComment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nickname := params["nickname"]
	PostID:= params["PostID"]

	body, _ := ioutil.ReadAll(r.Body)

	comment := &models.Comment{
		Comment_author: nickname,
		Comment_post:PostID,

	}
	if err := json.Unmarshal(body, comment); err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queries.CreateComment(comment)
	j, _ := json.Marshal(comment)
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleCommentGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ForumID := params["PostID"]

	result := queries.QueriesGetAllComment(ForumID)
	resp, _ := json.Marshal(result)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

