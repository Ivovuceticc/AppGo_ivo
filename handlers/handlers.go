package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Ivovuceticc/microblogging/middlew"
	"github.com/Ivovuceticc/microblogging/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*-En caso de que sólo especifiquemos, por ejemplo, 127.0.0.1:8080, el servidor sólo respondería a peticiones que se hagan en la computadora que es al mismo tiempo el servidor, no desde otras.
  -Si queremos que cualquier pc se conecte sabiendo mi ip, debo poner "":8080" o cualquier otro puerto que se pueda usar. */

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
