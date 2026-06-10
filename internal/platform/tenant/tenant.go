package tenant

import "context"

type contextKey string

const tenantKey contextKey = "tenant_id"

// WithTenant returns a copy of ctx carrying the tenant ID.
func WithTenant(ctx context.Context, tenantId string) context.Context {
	return context.WithValue(ctx, tenantKey, tenantId)
}

// FromContext extracts the tenant ID, returning ("", false) if absent or empty.
func FromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(tenantKey).(string)
	return id, ok && id != ""
}
