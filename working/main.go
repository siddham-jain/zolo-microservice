package main

import (
	"net/http"
	"fmt"
)
func main(){
	server := &http.Server{
		Addr: ":5050",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe();

	if err != nil {
		fmt.Println("failed to listen and serve", err);
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World")); 
}