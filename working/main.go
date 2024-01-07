package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
func main(){

	router := chi.NewRouter();
	router.Get("/hello", basicHandler);

	server := &http.Server{
		Addr: ":5050",
		Handler: router,
	}

	err := server.ListenAndServe();

	if err != nil {
		fmt.Println("failed to listen and serve", err);
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World")); 
}