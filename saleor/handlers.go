package saleor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/grzegorzderdak/saleor-wishlist/config"
)

func ManifestHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("ManifestHandler")

	config := config.Load()
	domain := config.AppDomain

	appManifest := Manifest{}
	appManifest.Initialize(
		"wishlist.app",
		"1.0.0",
		"Wishlist App",
		"Grzegorz Derdak",
		"Wishlist App is an Saleor app for adding products to customer wishlist.",
		domain+"/app",
		domain+"/configuration",
		domain+"/register",
		"",
		domain+"/app-data-privacy",
		domain+"/homepage",
		domain+"/support",
		[]Permission{ManageProducts, ManageUsers},
	)

	jsonResponseData, err := json.Marshal(appManifest)

	if err != nil {
		log.Fatal(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponseData)
	w.WriteHeader(http.StatusOK)
}

type RegisterPayload struct {
	AuthToken string `json:"auth_token"`
}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	var payload RegisterPayload

	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	fmt.Printf("AuthToken: %+v", payload.AuthToken)

	w.WriteHeader(http.StatusOK)
}

func AppHandler(w http.ResponseWriter, req *http.Request) {
	log.Default().Println("AppHandler")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func BaseHandler(w http.ResponseWriter, req *http.Request) {
	log.Default().Println("BaseHandler")

	w.WriteHeader(http.StatusOK)
}
