package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Taigore/ticket-go/azure-api/checkCode"
)

func getPort() string {
	val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if ok {
		return val
	} else {
		return "8080"
	}
}

func main() {
	port := getPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/checkCode", checkCode.Handle)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
