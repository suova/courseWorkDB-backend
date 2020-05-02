package controllers

import (
	"databasework/Sessions"
	"databasework/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func HandleLogOut(w http.ResponseWriter, r *http.Request) {
	println("in delete")
	body, _ := ioutil.ReadAll(r.Body)
	cookie:=&models.Session{}
	if err := json.Unmarshal(body, cookie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if Sessions.ExistinMap(cookie.Cookie){
		Sessions.DeletefromMap(cookie.Cookie)
		println("delete", cookie.Cookie)
		w.WriteHeader(http.StatusOK)
	}else{
		w.WriteHeader(http.StatusForbidden)
	}
	resp, _ := json.Marshal(nil)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleSendCookie(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	nickname := params["nickname"]
	cookie := Sessions.GetFormMap(nickname)
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(cookie)
	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}