package main

import "net/http"

func main(){

	http.ListenAndServe(":5050", nil)
}