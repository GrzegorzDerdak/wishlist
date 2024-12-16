package saleor

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

func (m *Manifest) Initialize(
	id string,
	version string,
	// requiredSaleorVersion string,
	name string,
	author string,
	about string,
	appUrl string,
	configurationUrl string,
	tokenTargetUrl string,
	dataPrivacy string,
	dataPrivacyUrl string,
	homepageUrl string,
	supportUrl string,
	permissions []Permission,
	// brand map[string]interface{},
	// extensions []map[string]interface{},
	// webhooks []map[string]interface{},

) Manifest {
	m.Id = id
	m.Version = version
	m.Name = name
	m.Author = author
	m.About = about
	m.AppUrl = appUrl
	m.ConfigurationUrl = configurationUrl
	m.TokenTargetUrl = tokenTargetUrl
	m.DataPrivacy = dataPrivacy
	m.DataPrivacyUrl = dataPrivacyUrl
	m.HomepageUrl = homepageUrl
	m.SupportUrl = supportUrl
	m.Permissions = permissions

	return *m
}
