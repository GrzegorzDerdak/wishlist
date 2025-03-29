package saleor

import (
	"encoding/json"
	"log"
	"net/http"
	"wishlist/internal"
)

// Permissions
type Permission string

// Saleor APP permissions
const (
	HandleCheckouts            Permission = "HANDLE_CHECKOUTS"
	HandlePayments             Permission = "HANDLE_PAYMENTS"
	HandleTaxes                Permission = "HANDLE_TAXES"
	ImpersonateUser            Permission = "IMPERSONATE_USER"
	ManageApps                 Permission = "MANAGE_APPS"
	ManageChannels             Permission = "MANAGE_CHANNELS"
	ManageCheckouts            Permission = "MANAGE_CHECKOUTS"
	ManageDiscounts            Permission = "MANAGE_DISCOUNTS"
	ManageGiftCard             Permission = "MANAGE_GIFT_CARD"
	ManageMenus                Permission = "MANAGE_MENUS"
	ManageObservability        Permission = "MANAGE_OBSERVABILITY"
	ManageOrders               Permission = "MANAGE_ORDERS"
	ManageOrdersImport         Permission = "MANAGE_ORDERS_IMPORT"
	ManagePages                Permission = "MANAGE_PAGES"
	ManagePageTypesAndAttrs    Permission = "MANAGE_PAGE_TYPES_AND_ATTRIBUTES"
	ManagePlugins              Permission = "MANAGE_PLUGINS"
	ManageProducts             Permission = "MANAGE_PRODUCTS"
	ManageProductTypesAndAttrs Permission = "MANAGE_PRODUCT_TYPES_AND_ATTRIBUTES"
	ManageSettings             Permission = "MANAGE_SETTINGS"
	ManageShipping             Permission = "MANAGE_SHIPPING"
	ManageStaff                Permission = "MANAGE_STAFF"
	ManageTaxes                Permission = "MANAGE_TAXES"
	ManageTranslations         Permission = "MANAGE_TRANSLATIONS"
	ManageUsers                Permission = "MANAGE_USERS"
)

type Manifest struct {
	About            string       `json:"about"`            // Description of the app
	AppUrl           string       `json:"appUrl"`           // URL of the app
	Author           string       `json:"author"`           // Author of the app
	ConfigurationUrl string       `json:"configurationUrl"` // URL of the configuration
	DataPrivacy      string       `json:"dataPrivacy"`      // Data privacy
	DataPrivacyUrl   string       `json:"dataPrivacyUrl"`   // URL of the data privacy
	HomepageUrl      string       `json:"homepageUrl"`      // URL of the homepage
	Id               string       `json:"id"`               // ID of the app
	Name             string       `json:"name"`             // Name of the app
	SupportUrl       string       `json:"supportUrl"`       // URL of the support
	TokenTargetUrl   string       `json:"tokenTargetUrl"`   // URL of the token target
	Version          string       `json:"version"`          // Version of the app
	Permissions      []Permission `json:"permissions"`      // List of permissions required by the app
}
type RegisterPayload struct {
	AuthToken string `json:"auth_token"`
}

type SaleorManifestHandler struct {
	SaleorManifestService *SaleorManifestService
}

func NewSaleorManifestHandler(saleorManifestService *SaleorManifestService) *SaleorManifestHandler {
	return &SaleorManifestHandler{
		SaleorManifestService: saleorManifestService,
	}
}

func (s *SaleorManifestHandler) ManifestGetHandler(w http.ResponseWriter, r *http.Request) {
	config := internal.NewConfig()
	domain := config.AppDomain

	appManifest := &Manifest{
		Id:               "wishlist.app",
		Name:             "Wishlist app",
		Version:          "1.0.0",
		DataPrivacy:      "",
		About:            "A dead simple wishlist app for Saleor.",
		AppUrl:           domain + "/saleor/app",
		Author:           "Grzegorz Derdak <grzegorz.derdak@gmail.com>",
		ConfigurationUrl: domain + "/saleor/configuration",
		DataPrivacyUrl:   domain + "/saleor/app-data-privacy",
		HomepageUrl:      domain + "/saleor/homepage",
		SupportUrl:       domain + "/saleor/support",
		Permissions:      []Permission{ManageProducts, ManageUsers},
		TokenTargetUrl:   domain + "/saleor/register",
	}

	jsonResponseData, err := json.Marshal(appManifest)

	if err != nil {
		log.Fatal(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponseData)
}

func (s *SaleorManifestHandler) ManifestRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RegisterPayload
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	headers := ParseSaleorHeaders(r)
	if err := headers.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	saleorConfig := &SaleorConfig{
		Domain:        headers.Domain,
		ApiUrl:        headers.ApiUrl,
		SchemaVersion: headers.SchemaVersion,
		AuthToken:     data.AuthToken,
	}

	if _, err := s.SaleorManifestService.RegisterSaleorDomain(saleorConfig); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
