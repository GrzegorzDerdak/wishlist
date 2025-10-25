package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"wishlist/internal"
	"wishlist/logger"
	"wishlist/saleor"
	"wishlist/wishlists"
)

func main() {
	l := logger.Get()
	config := internal.NewConfig()

	db, err := internal.ConnectToDatabase(config.DSN)
	if err != nil {
		l.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	db.AutoMigrate(&wishlists.Item{})
	db.AutoMigrate(&wishlists.Wishlist{})
	db.AutoMigrate(&saleor.SaleorConfig{})

	r := mux.NewRouter().StrictSlash(false)
	r.Use(internal.RequestIDMiddleware)

	// Healthcheck endpoint
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Saleor API
	saleorApi := r.PathPrefix("/saleor").Subrouter()
	saleorHandler := saleor.NewSaleorHandler(
		saleor.NewSaleorManifestService(
			saleor.NewSaleorRepository(db),
		),
	)
	saleorHandler.RegisterRoutes(saleorApi)

	authMiddleware, err := internal.CreateAuthMiddleware(config.JWKSURL)

	if err != nil {
		l.Fatal().Err(err).Msg("Could not create auth middleware.")
	}

	// Wishlists API
	wishlistsApi := r.PathPrefix("/api/v1/wishlists").Subrouter()
	wishlistsApi.Use(authMiddleware.Middleware)

	wishlistHandler := wishlists.NewWishlistHandler(
		wishlists.NewWishlistService(
			wishlists.NewWishlistRepository(db),
		),
	)
	wishlistHandler.RegisterRoutes(wishlistsApi)

	// Default handler
	http.Handle("/", r)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: r,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	l.Debug().Str("Starting server on port", config.Port).Msg("Server starting")
	server.ListenAndServe()
}
