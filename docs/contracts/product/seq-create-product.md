# Sequence — POST /products

    Merchant -> Web App        : fill form, submit
    Web App  -> Marketplace API: POST /products (X-Tenant-ID, JSON body)
    API      -> tenant         : FromContext(ctx) -> tenant_id (else 400)
    API      -> domain         : NewProduct(tenant_id, name, price, currency, category)
    alt invalid
        domain --> API        : error (e.g. ErrEmptyName)
        API    --> Web App     : 422 { error, field }
    else valid
        API    -> Repository   : Create(ctx, product)
        Repository -> Postgres : INSERT ... RETURNING id, created_at, updated_at
        Postgres   --> Repository : row
        Repository -> domain   : Reconstitute(row...) -> *Product
        Repository --> API     : *Product
        API        --> Web App : 201 { product }
    end