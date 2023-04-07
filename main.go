package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gevs90/crypto-api/api/routes"
	"github.com/gevs90/crypto-api/config"
)

func main() {
	config.LoadAppConfig()
	routes := routes.NewRouter()
	var port string = config.AppConfig.Port
	server := http.ListenAndServe(port, routes)

	log.Fatal(server)
	fmt.Println("El servidor esta corriendo en http://localhost" + port)

}
