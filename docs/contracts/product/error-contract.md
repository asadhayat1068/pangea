# Error Contract (v1)

All errors return a JSON body of the shape:
    { "error": "<human message>", "field": "<offending field or empty>" }

Mapping from domain validation errors to HTTP:
| Domain error (Go)   | Message                                | field        | HTTP Code |
|---------------------|----------------------------------------|--------------|------|
| ErrMissingTenant    | tenant_id is required                  | tenant_id    | 400  |
| ErrEmptyName        | name is required                       | name         | 422  |
| ErrNameTooLong      | name must be at most 140 characters    | name         | 422  |
| ErrNegativePrice    | price_cents must be >= 0               | price_cents  | 422  |
| ErrEmptyCategory    | category is required                   | category     | 422  |
| (malformed JSON)    | invalid request body                   | (empty)      | 400  |

Note: a missing tenant is a 400 (the caller failed to supply context), while a
bad field value is a 422 (the body was understood but is semantically invalid).