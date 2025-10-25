package saleor

import (
	"encoding/json"
	"net/http"
	"wishlist/internal"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
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

type EventName string

// Saleor webhook events
const (
	CheckoutCreated                EventName = "CHECKOUT_CREATED"
	CheckoutFilterShippingMethods  EventName = "CHECKOUT_FILTER_SHIPPING_METHODS"
	CheckoutUpdated                EventName = "CHECKOUT_UPDATED"
	CustomerCreated                EventName = "CUSTOMER_CREATED"
	CustomerUpdated                EventName = "CUSTOMER_UPDATED"
	DraftOrderCreated              EventName = "DRAFT_ORDER_CREATED"
	DraftOrderDeleted              EventName = "DRAFT_ORDER_DELETED"
	DraftOrderUpdated              EventName = "DRAFT_ORDER_UPDATED"
	FulfillmentCanceled            EventName = "FULFILLMENT_CANCELED"
	FulfillmentCreated             EventName = "FULFILLMENT_CREATED"
	InvoiceDeleted                 EventName = "INVOICE_DELETED"
	InvoiceRequested               EventName = "INVOICE_REQUESTED"
	InvoiceSent                    EventName = "INVOICE_SENT"
	NotifyUser                     EventName = "NOTIFY_USER"
	OrderCancelled                 EventName = "ORDER_CANCELLED"
	OrderConfirmed                 EventName = "ORDER_CONFIRMED"
	OrderCreated                   EventName = "ORDER_CREATED"
	OrderFilterShippingMethods     EventName = "ORDER_FILTER_SHIPPING_METHODS"
	OrderFulfilled                 EventName = "ORDER_FULFILLED"
	OrderFullyPaid                 EventName = "ORDER_FULLY_PAID"
	OrderUpdated                   EventName = "ORDER_UPDATED"
	PageCreated                    EventName = "PAGE_CREATED"
	PageDeleted                    EventName = "PAGE_DELETED"
	PageUpdated                    EventName = "PAGE_UPDATED"
	PaymentAuthorize               EventName = "PAYMENT_AUTHORIZE"
	PaymentCapture                 EventName = "PAYMENT_CAPTURE"
	PaymentConfirm                 EventName = "PAYMENT_CONFIRM"
	PaymentListGateways            EventName = "PAYMENT_LIST_GATEWAYS"
	PaymentProcess                 EventName = "PAYMENT_PROCESS"
	PaymentRefund                  EventName = "PAYMENT_REFUND"
	PaymentVoid                    EventName = "PAYMENT_VOID"
	ProductCreated                 EventName = "PRODUCT_CREATED"
	ProductDeleted                 EventName = "PRODUCT_DELETED"
	ProductUpdated                 EventName = "PRODUCT_UPDATED"
	ProductVariantBackInStock      EventName = "PRODUCT_VARIANT_BACK_IN_STOCK"
	ProductVariantCreated          EventName = "PRODUCT_VARIANT_CREATED"
	ProductVariantDeleted          EventName = "PRODUCT_VARIANT_DELETED"
	ProductVariantOutOfStock       EventName = "PRODUCT_VARIANT_OUT_OF_STOCK"
	ProductVariantUpdated          EventName = "PRODUCT_VARIANT_UPDATED"
	SaleCreated                    EventName = "SALE_CREATED"
	SaleDeleted                    EventName = "SALE_DELETED"
	SaleUpdated                    EventName = "SALE_UPDATED"
	ShippingListMethodsForCheckout EventName = "SHIPPING_LIST_METHODS_FOR_CHECKOUT"
	TranslationCreated             EventName = "TRANSLATION_CREATED"
	TranslationUpdated             EventName = "TRANSLATION_UPDATED"
)

type Webhook struct {
	Name        string      `json:"name"`        // Name of the webhook
	AsyncEvents []EventName `json:"asyncEvents"` // List of asynchronous events
	Query       string      `json:"query"`       // GraphQL subscription query
	TargetUrl   string      `json:"targetUrl"`   // Target URL for the webhook
	IsActive    bool        `json:"isActive"`    // Is the webhook active
}

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
	Webhooks         []Webhook    `json:"webhooks"`         // List of webhooks
}
type RegisterPayload struct {
	AuthToken string `json:"auth_token"`
}

type SaleorHandler struct {
	SaleorService *SaleorService
}

func NewSaleorHandler(saleorService *SaleorService) *SaleorHandler {
	return &SaleorHandler{
		SaleorService: saleorService,
	}
}

func (s *SaleorHandler) ManifestGetHandler(w http.ResponseWriter, r *http.Request) {
	config := internal.NewConfig()
	domain := config.AppDomain
	l := zerolog.Ctx(r.Context())

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
		Webhooks: []Webhook{
			{
				Name:        "Product Events",
				AsyncEvents: []EventName{ProductCreated, ProductUpdated, ProductDeleted, ProductVariantBackInStock, ProductVariantOutOfStock},
				Query:       "subscription { event { ... on ProductCreated { product { id } } ... on ProductUpdated { product { id } } ... on ProductDeleted { product { id } } ... on ProductVariantBackInStock { productVariant { id } } ... on ProductVariantOutOfStock { productVariant { id } } } }",
				TargetUrl:   domain + "/saleor/webhooks/product-events",
				IsActive:    true,
			},
		},
	}

	jsonResponseData, err := json.Marshal(appManifest)

	if err != nil {
		l.Fatal().Err(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(jsonResponseData)
}

func (s *SaleorHandler) ManifestRegisterHandler(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())

	if r.Method != http.MethodPost {
		l.Warn().Str("method", r.Method).Msg("Invalid method for manifest register handler")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RegisterPayload
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		l.Warn().Msg("Failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	headers := ParseSaleorHeaders(r)
	if err := headers.Validate(); err != nil {
		l.Warn().Msg("Failed to validate Saleor headers")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	saleorConfig := &SaleorConfig{
		Domain:        headers.Domain,
		ApiUrl:        headers.ApiUrl,
		SchemaVersion: headers.SchemaVersion,
		AuthToken:     data.AuthToken,
	}

	if _, err := s.SaleorService.RegisterSaleorDomain(saleorConfig); err != nil {
		l.Err(err).Msg("Failed to register Saleor domain")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SaleorHandler) ProductEventsWebhookHandler(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())

	if r.Method != http.MethodPost {
		l.Warn().Str("method", r.Method).Msg("Invalid method for product events webhook")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	headers := ParseSaleorHeaders(r)
	if err := headers.Validate(); err != nil {
		l.Warn().Msg("Failed to validate Saleor headers")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		l.Warn().Msg("Failed to decode webhook payload")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	l.Debug().Interface("payload", payload).Msg("Received product event webhook")
	w.WriteHeader(http.StatusOK)
}

func (s *SaleorHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/manifest", s.ManifestGetHandler)
	// r.HandleFunc("/app", saleor.AppHandler)
	r.HandleFunc("/register", s.ManifestRegisterHandler)
	r.HandleFunc("/webhooks/product-events", s.ProductEventsWebhookHandler)
	// r.HandleFunc("/support", saleor.BaseHandler)
	// r.HandleFunc("/homepage", saleor.BaseHandler)
	// r.HandleFunc("/configuration", saleor.BaseHandler)
}
