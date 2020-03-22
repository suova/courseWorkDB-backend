package controllers

import (
	"databasework/models"
	"databasework/queries"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nickname := params["nickname"]
	ForumTitle := params["ForumID"]

	body, _ := ioutil.ReadAll(r.Body)

	post := &models.Post{
		Post_author: nickname,
		Post_thread : ForumTitle,
	}
	if err := json.Unmarshal(body, post); err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queries.CreatePost(post)
	j, _ := json.Marshal(post)
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func HandlePostGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ForumID := params["ForumID"]

	result := queries.QueriesGetAllPost(ForumID)
	resp, _ := json.Marshal(result)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func HandleOnePostGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ForumID := params["PostID"]

	result := queries.QueriesGetOnePost(ForumID)
	resp, _ := json.Marshal(result)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func HandleOnePostDelete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ForumID := params["PostID"]

	result := queries.QueriesDeleteOnePost(ForumID)
	if result{
		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusForbidden)
		return
	}

}
