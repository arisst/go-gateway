package main

import (
	"fmt"
	"go-gateway/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func appRoute() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/sendsms", controllers.SendSMS).Methods("GET")

	return router
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	fmt.Println(port)
	if port == "" {
		port = "8001"
	}

	fmt.Println("app running at: http://localhost:" + port)

	err := http.ListenAndServe(":"+port, appRoute())
	if err != nil {
		fmt.Print(err)
	}

}
