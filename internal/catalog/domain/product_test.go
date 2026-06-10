package domain

import (
	"errors"
	"testing"
)

func TestNewProduct_Valid(t *testing.T) {
	p, err := NewProduct(
		"tenant-freshcart",
		"Organic Bananas",
		int64(199),
		"USD",
		"Produce",
	)

	if err != nil {
		t.Fatalf("expected: no error, got: %v", err)
	}

	if p.TenantID != "tenant-freshcart" {
		t.Errorf("tenant_id not set, got %q", p.TenantID)
	}

	if p.InStock {
		t.Errorf("a new product should default to out of stock")
	}

	if p.Currency != "USD" {
		t.Errorf("expected currency USD, got %q", p.Currency)
	}
}

func TestNewProduct_Invalid(t *testing.T) {
	cases := []struct {
		label       string
		tenantId    string
		productName string
		price       int64
		category    string
		wantErr     error
	}{
		{"missing tenant", "", "Milk", 100, "Dairy", ErrMissingTenant},
		{"empty name", "t1", "", 100, "Dairy", ErrEmptyName},
		{"negative price", "t2", "Milk", -100, "Dairy", ErrNegativePrice},
		{"empty cateory", "t3", "Tomatoes", 150, "", ErrEmptyCategory},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			_, err := NewProduct(tc.tenantId, tc.productName, tc.price, "USD", tc.category)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("expected: %v, got: %v", tc.wantErr, err)
			}

		})
	}
}
