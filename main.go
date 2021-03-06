package main

import (
	"databasework/controllers"
	"databasework/queries"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	queries.InitDB()
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{
			" http://localhost:8080/"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	r.HandleFunc("/",controllers.HandleThreadGet).Methods("GET")
	r.HandleFunc("/forum/create",controllers.HandleCreateThread).Methods("POST")

	r.HandleFunc("/user/{nickname}/create", controllers.HandleUserPost).Methods("POST")
	r.HandleFunc("/user/{nickname}/profile", controllers.HandleUserGet).Methods("GET")
	r.HandleFunc("/getUsers", controllers.HandleUsersGet).Methods("GET")
	r.HandleFunc("/changeRole/{nickname}", controllers.HandleChangeRole).Methods("POST")
	r.HandleFunc("/user/{nickname}/singIn", controllers.HandleUsersignin).Methods("POST")

	r.HandleFunc("/forum/{ForumID}/post/{nickname}/create",controllers.HandleCreatePost).Methods("POST")
	r.HandleFunc("/forum/{ForumID}/posts",controllers.HandlePostGet).Methods("GET")

	r.HandleFunc("/post/{PostID}/comment/{nickname}/create",controllers.HandleCreateComment).Methods("POST")
	r.HandleFunc("/post/{PostID}/comments",controllers.HandleCommentGet).Methods("GET")
	r.HandleFunc("/post/{PostID}",controllers.HandleOnePostGet).Methods("GET")
	r.HandleFunc("/post/{PostID}/delete",controllers.HandleOnePostDelete).Methods("GET")

	r.HandleFunc("/like/{CommentID}/{nickname}",controllers.HandleLike).Methods("GET")
	r.HandleFunc("/dislike/{CommentID}/{nickname}",controllers.HandleDislike).Methods("GET")

	r.HandleFunc("/logout",controllers.HandleLogOut).Methods("POST")
	r.HandleFunc("/sendCookie/{nickname}",controllers.HandleSendCookie).Methods("GET")

	log.Fatal(http.ListenAndServe(":8181", corsMiddleware(r)))
	//srv := &http.Server{
	//	Handler: r,
	//	Addr:    "127.0.0.1:8181",
	//	// Good practice: enforce timeouts for servers you create!
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//log.Fatal(srv.ListenAndServe())

}