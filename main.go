package main

import (
	"fmt"
	"log"
	"net/http"
	"tui.com/API/routes"
)

func main(){
	port := ":3000"
	router := routes.CreateRouter()
	fmt.Println("serving on" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
