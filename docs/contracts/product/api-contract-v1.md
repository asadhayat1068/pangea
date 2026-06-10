# Marketplace API Contract (v1)

## Tenancy

Until the auth sprint, the tenant is carried in the request 

```
header:
    X-Tenant-ID: <uuid>
```
The server reads it and injects it into the request context (ADR-002). The client never puts tenant_id in the body. Later this is replaced by a tenant claim derived from the authenticated user's token.

## POST /products — create a product

***Request headers:***
```
    Content-Type: application/json
    X-Tenant-ID: <uuid>
```
***Request body:***
```
    {
      "name": "Organic Bananas",
      "price_cents": 199,
      "currency": "USD",        // optional, defaults to "USD"
      "category": "Produce",
      "in_stock": false          // optional, defaults to false
    }
```
***Responses:***

```    
    {
    201 Created
      "product": {
        "id": "uuid",
        "tenant_id": "uuid",
        "name": "Organic Bananas",
        "price_cents": 199,
        "currency": "USD",
        "category": "Produce",
        "in_stock": false,
        "created_at": "RFC3339 timestamp",
        "updated_at": "RFC3339 timestamp"
      }
    }
```
    422 Unprocessable Entity  -> validation failure (see error contract)


    400 Bad Request           -> malformed JSON or missing X-Tenant-ID

## GET /products — list this tenant's products

***Request headers:***
```
    X-Tenant-ID: <uuid>
```
***Responses:***
``` 
    200 OK
    { "products": [ { ...product... }, ... ] }   // only this
```    
tenant's products

    400 Bad Request -> missing X-Tenant-ID