# PRD-001 — Create & List a Product

## Problem
FreshCart has no system of record for its products. We will give merchants one.

## Target user
FreshCart merchant (Dana, Marco).

## In-scope user story (exactly one)
As a FreshCart merchant, I can create a product so customers can later buy it.

## Success metric
A merchant creates a valid product in under 60 seconds and immediately sees it
in the product list.

## Acceptance criteria (handed to the BA to sharpen)
- A product can be created with: name, price, category (stock flag optional,
  defaults to "out of stock").
- Invalid products are rejected with a clear, specific error.
- Every product is permanently associated with a tenant (`tenant_id`).
- Created products can be listed back, scoped to the tenant.

## Acceptance Scenarios (Given/When/Then) 

> These are the basis for the sprint's Demo Scenario, and the BA should add more as needed to cover edge cases and error conditions. The more scenarios, the better — they are the basis for our automated tests and documentation.

### Scenario: Create a valid product
Given I am a FreshCart merchant
When I submit a product with name "Organic Bananas", price 199 cents, category "Produce"
Then the product is created
And the response has status 201
And the response body contains a generated id and the tenant_id
And in_stock defaults to false and currency defaults to "USD"

### Scenario: Reject a blank name
Given I am a FreshCart merchant
When I submit a product with an empty name
Then no product is created
And the response has status 422
And the response body is { "error": "name is required", "field": "name" }

### Scenario: Reject a negative price
When I submit a product with price -1
Then the response has status 422
And the response body is { "error": "price_cents must be >= 0", "field": "price_cents" }

### Scenario: Reject a missing category
When I submit a product without a category
Then the response has status 422
And the response body is { "error": "category is required", "field": "category" }

### Scenario: Missing tenant
Given the request has no tenant context
When I submit any product
Then the response has status 400
And the response body is { "error": "tenant_id is required", "field": "tenant_id" }

### Scenario: List products is tenant-scoped
Given tenant A has 2 products and tenant B has 1 product
When tenant A lists products
Then exactly tenant A's 2 products are returned, and never tenant B's


## Explicitly out of scope for this PRD
Cart, checkout, payments, images, search, customer accounts.

## Dependencies
None — this is the first feature. It depends only on the Sprint-1 skeleton.