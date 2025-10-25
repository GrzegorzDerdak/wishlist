# 🎁 Wishlist App for Saleor

A powerful, installable Saleor extension that brings wishlist functionality to your e-commerce store. Let your customers save products they love and come back to them later—boosting engagement and conversions.

## ✨ Why Wishlist?

- **Increase Customer Retention** - Keep shoppers engaged by letting them save items for later
- **Boost Conversions** - Wishlists act as a gentle reminder, bringing customers back to complete purchases
- **Understand Your Customers** - Track which products customers are interested in
- **Social Proof** - Enable wishlist sharing to drive organic traffic

## 🚀 Features

### Core Functionality ✅

- [x] **Saleor Integration** - Seamless installation via `/manifest` endpoint
- [x] **Wishlist Management**
  - [x] Create new wishlists
  - [x] Retrieve single or all wishlists per user
  - [x] Update wishlist details (rename, modify)
  - [x] Delete wishlists
- [ ] **Product Management**
  - [ ] Add products to wishlist
  - [ ] Remove products from wishlist
- [ ] **Webhook Support** - Real-time synchronization with Saleor events
- [x] **Authentication** - Secure access with JWT and JWKS
- [x] **Health Check Endpoint** - Monitor service status easily

### 🎯 Planned Enhancements

- [ ] **Social Sharing** - Share wishlists via link or social media
- [ ] **E-mail Notifications** - Get alerted when wishlist items are back in stock
- [ ] **Quick Actions** - Add items to cart directly from wishlist

## 🛠️ Tech Stack

- **Backend**: Go 1.25+ with Gorilla Mux
- **Database**: PostgreSQL 16
- **Authentication**: JWT with JWKS
- **Logging**: Structured logging with Zerolog
- **Container**: Docker & Docker Compose ready
- **Cloud**: AWS ECS compatible

## 🏃 Quick Start

### Prerequisites

- Go 1.25 or higher
- PostgreSQL 16
- Docker & Docker Compose (optional)

### Local Development

1. **Clone the repository**

```bash
   git clone https://github.com/GrzegorzDerdak/wishlist
   cd wishlist-
```

2. **Set up environment variables**

```bash
   cp .env.example .env
   # Edit .env with your configuration
```

3. **Run with Docker Compose**

```bash
   docker-compose up --build
```

4. **Or run locally**

```bash
   go mod download
   go run main.go
```

The API will be available at `http://localhost:8080`

### Installation in Saleor

1. In your Saleor dashboard, go to **Apps** → **Install External App**
2. Enter the manifest URL: `http://your-domain:8080/manifest`
3. Authorize the app
4. Start using wishlists! 🎉

## 📚 API Documentation

### Endpoints

**Wishlists**

- `GET /api/v1/wishlists` - Get all wishlists for authenticated user
- `POST /api/v1/wishlists` - Create a new wishlist
- `GET /api/v1/wishlists/{id}` - Get specific wishlist
- `PUT /api/v1/wishlists/{id}` - Update wishlist
- `DELETE /api/v1/wishlists/{id}` - Delete wishlist

**Health Check**

- `GET /healthcheck` - Service health status

## 🏗️ Architecture

```txt
wishlist-app/
├── bruno/                 # Bruno API collection for testing
├── internal/              # Private application code
│   ├── config.go           # Configuration management
│   ├── middleware.go      # Request middleware
│   └── database.go        # Database connection
├── wishlists/             # Wishlist domain
│   ├── handler.go         # HTTP handlers
│   ├── service.go         # Business logic
│   ├── repository.go      # Data access
│   └── models.go          # Domain models
├── saleor/                # Saleor integration
│   ├── auth.go            # HTTP handlers
│   ├── handlers.go        # Saleor manifest and webhooks
│   ├── repository.go      # Data access
│   └── models.go          # Domain models
├── logger/                # Logging utilities
└── docker-compose.yml     # Container orchestration
└── Dockerfile              # Docker image definition
```

## 📄 License

[LICENSE](LICENSE) © 2025 Grzegorz Derdak

## 🔗 Links

- [Saleor Documentation](https://docs.saleor.io/)
- [Report Issues](https://github.com/your-repo/issues)
