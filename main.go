package main

import (
	"log"
	"net/http"

	"github.com/balagrivine/linky/config"
	"github.com/balagrivine/linky/logic"

	"github.com/joho/godotenv"
)

func main() {

	// Load environmental variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Unble to locate .env file")
	}

	// Initialize the configuration
	dbCfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Couldn't initialize database configuration: ", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users", logic.HandleCreateUser(dbCfg))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
