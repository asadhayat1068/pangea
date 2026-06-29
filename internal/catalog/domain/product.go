package domain

import (
	"errors"
	"strings"
	"time"
)

// Validation errors returned by the Product constructor. They mirror the
// error strings in docs/product-data-dictionary.md exactly.
var (
	ErrMissingTenant = errors.New("tenant_id is required")
	ErrEmptyName     = errors.New("name is required")
	ErrNameTooLong   = errors.New("name must be at most 140 characters")
	ErrNegativePrice = errors.New("price_cents must be >= 0")
	ErrEmptyCategory = errors.New("category is required")
)

const maxNameLength = 140
const defaultCurrency = "USD"

// Product is the core catalog entity. It is always scoped to a tenant.
// Invariants live HERE, in the domain — never in the HTTP layer (ADR-001/002).
type Product struct {
	ID         string
	TenantID   string
	Name       string
	PriceCents int64
	Currency   string
	Category   string
	InStock    bool
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

// NewProduct constructs a valid Product or returns an error explaining why the
// input is invalid. ID is assigned later by the persistence adapter.

func NewProduct(tenantID, name string, priceCents int64, currency, category string) (*Product, error) {
	tenantID = strings.TrimSpace(tenantID)
	if tenantID == "" {
		return nil, ErrMissingTenant
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrEmptyName
	}

	if len(name) > maxNameLength {
		return nil, ErrNameTooLong
	}

	if priceCents < 0 {
		return nil, ErrNegativePrice
	}

	currency = strings.TrimSpace(currency)
	if currency == "" {
		currency = defaultCurrency
	}

	category = strings.TrimSpace(category)
	if category == "" {
		return nil, ErrEmptyCategory
	}

	now := time.Now().UTC()

	return &Product{
		TenantID:   tenantID,
		Name:       name,
		PriceCents: priceCents,
		Currency:   currency,
		Category:   category,
		InStock:    false,
		UpdatedAt:  now,
		CreatedAt:  now,
	}, nil
}

func Reconstitute(id, tenantID, name string, priceCents int64, currency, category string, inStock bool, createdAt, updatedAt time.Time) *Product {
	return &Product{
		ID:         id,
		TenantID:   tenantID,
		Name:       name,
		PriceCents: priceCents,
		Currency:   currency,
		Category:   category,
		InStock:    inStock,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}
