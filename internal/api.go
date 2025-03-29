package internal

import (
	"net/http"
)

// Placeholder handler functions
// func createWishlist(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("createWishlist"))
// }

// func updateWishlist(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("updateWishlist"))
// }

// func deleteWishlist(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("deleteWishlist"))
// }

func HandleGetWishlistByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("getWishlist"))
}

// func listWishlists(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("listWishlists"))
// }

// func addItem(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("addItem"))
// }

// func removeItem(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("removeItem"))
// }

// func updateItem(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("updateItem"))
// }

// func listItems(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("listItems"))
// }

// func publishWishlist(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("publishWishlist"))
// }

// func unpublishWishlist(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("unpublishWishlist"))
// }

func HandleNotImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusExpectationFailed)
	w.Write([]byte("NOT_IMPLEMENTED"))
}
