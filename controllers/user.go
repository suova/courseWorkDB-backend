package controllers

import (
	"databasework/Sessions"
	"databasework/models"
	"databasework/queries"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func HandleUserPost(w http.ResponseWriter, r *http.Request) {
	cookie := Sessions.String(50)

	params := mux.Vars(r)
	nickname := params["nickname"]
	Sessions.AddtoMap(cookie, nickname)


	user:=&models.User{
		Nickname: nickname,
	}
	if(nickname=="admin"){
		user.Role = 3
	}
	if(nickname=="moderator"){
		user.Role = 2
	}
	body, _ := ioutil.ReadAll(r.Body)


	if err := json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result:=queries.User(user)
	if result == "OK" {
		j, _ := json.Marshal(user)
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write(j); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}else{
		w.WriteHeader(http.StatusConflict)
		j:= `exist nickname`
		if _, err := w.Write([]byte(j)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}


}

func HandleUserGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nickname := params["nickname"]
	result := queries.FindUser(nickname)

	resp, _ := json.Marshal(result)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleUsersGet(w http.ResponseWriter, r *http.Request) {

	result := queries.FindUsers()

	resp, _ := json.Marshal(result)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleChangeRole(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nickname := params["nickname"]
	body, _ := ioutil.ReadAll(r.Body)
	role:=&models.Role{}
	if err := json.Unmarshal(body, role); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := queries.ChangeRole(nickname,role.Role)
	resp, _ := json.Marshal(err)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleUsersignin(w http.ResponseWriter, r *http.Request) {
	cookie := Sessions.String(50)

	params := mux.Vars(r)
	nickname := params["nickname"]
	Sessions.AddtoMap(cookie, nickname)

	body, _ := ioutil.ReadAll(r.Body)
	user:=&models.User{}
	if err := json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := queries.FindUser(nickname)

	if result.Password == user.Password {
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(result)
		if _, err := w.Write(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}else{
		w.WriteHeader(http.StatusForbidden)
	}
}

