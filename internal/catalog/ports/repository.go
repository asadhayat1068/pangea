package ports

import (
	"context"

	"github.com/asadhayat1068/pangea/internal/catalog/domain"
)

// ProductRepository is the OUTBOUND port for persisting catalog products.
// The Postgres adapter (built on Days 4–6) will implement this interface.

type ProductRepository interface {
	Create(ctx context.Context, p *domain.Product) (*domain.Product, error)
	ListByTenants(ctx context.Context, tenantID string) ([]*domain.Product, error)
}
