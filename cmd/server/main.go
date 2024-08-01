package main

import (
	"log"

	"github.com/surajjyoti-personal/proglog/internal/server"
)

func main() {
	httpServ := server.NewHTTPServer(":8080")
	log.Fatal(httpServ.ListenAndServe())
}
