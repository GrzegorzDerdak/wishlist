package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"wishlist/internal"
	"wishlist/saleor"
)

func main() {
	config := internal.NewConfig()
	db, err := internal.ConnectToDatabase(config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&internal.Item{})
	db.AutoMigrate(&internal.Wishlist{})
	db.AutoMigrate(&saleor.SaleorConfig{})

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	wishlistHandler := internal.NewWishlistHandler(internal.NewWishlistService(internal.NewWishlistRepository(db)))
	// Temporary handler for unimplemented endpoints
	protectedRoute := internal.AuthMiddleware(http.HandlerFunc(internal.HandleNotImplemented))

	// Wishlist management
	r.Handle("/api/v1/wishlists", internal.AuthMiddleware(http.HandlerFunc(wishlistHandler.Create))).Methods("POST")

	r.Handle("/api/v1/wishlists", protectedRoute).Methods("GET")
	r.Handle("/api/v1/wishlists/{wishlist_id}", protectedRoute).Methods("PUT")
	r.Handle("/api/v1/wishlists/{wishlist_id}", protectedRoute).Methods("DELETE")
	r.HandleFunc("/api/v1/wishlists/{wishlist_id}", wishlistHandler.GetByID).Methods("GET")

	// Wishlist item management
	r.Handle("/api/v1/wishlists/{wishlist_id}/items", protectedRoute).Methods("POST")
	r.Handle("/api/v1/wishlists/{wishlist_id}/items/{item_id}", protectedRoute).Methods("DELETE")
	r.Handle("/api/v1/wishlists/{wishlist_id}/items/{item_id}", protectedRoute).Methods("PUT")
	r.Handle("/api/v1/wishlists/{wishlist_id}/items", protectedRoute).Methods("GET")

	// Wishlist publishing
	r.Handle("/api/v1/wishlists/{wishlist_id}/publish", protectedRoute).Methods("POST")
	r.Handle("/api/v1/wishlists/{wishlist_id}/unpublish", protectedRoute).Methods("POST")

	// Saleor URLs
	saleorConfigRepository := saleor.NewSaleorConfigRepository(db)
	saleorConfigService := saleor.NewSaleorManifestService(saleorConfigRepository)
	saleorManifestHandler := saleor.NewSaleorManifestHandler(saleorConfigService)

	r.HandleFunc("/saleor/manifest", saleorManifestHandler.ManifestGetHandler)
	// r.HandleFunc("/saleor/app", saleor.AppHandler)
	r.HandleFunc("/saleor/register", saleorManifestHandler.ManifestRegisterHandler)
	// r.HandleFunc("/saleor/support", saleor.BaseHandler)
	// r.HandleFunc("/saleor/homepage", saleor.BaseHandler)
	// r.HandleFunc("/saleor/configuration", saleor.BaseHandler)

	// Default handler
	http.Handle("/", r)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: r,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	log.Println("Starting server on " + ":" + config.Port)
	server.ListenAndServe()
}
