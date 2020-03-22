package controllers

import (
	"databasework/models"
	"databasework/queries"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleLike(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nickname := params["nickname"]
	commentID := params["CommentID"]

	like :=&models.Like{
		Like_author:nickname,
		Like_comment:commentID,
	}

	queries.Like(like)
	j, _ := json.Marshal(like)
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleDislike(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nickname := params["nickname"]
	commentID := params["CommentID"]

	like :=&models.Like{
		Like_author:nickname,
		Like_comment:commentID,
	}

	queries.DisLike(like)
	j, _ := json.Marshal(like)
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
