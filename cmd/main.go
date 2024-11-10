package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shaik-baleeghuddin-kashif/auth-server/config"
)

func main() {
	// Parse Flags
	config.Flags()

	// Load Config
	config.LoadConfig()
	port := config.AppConfig.Port

	// Connect Database
	fmt.Println(config.Database)

	// SetUp Routes
	http.HandleFunc("/", homehandler)

	// Start Server
	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Auth Server")
}
