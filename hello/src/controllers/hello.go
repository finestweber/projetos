package controllers

import (
	"net/http"

)
func Hello (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Ola mundo"))
}