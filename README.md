#Note API

## Tech Stack
- Go
- Gin
- GORM
- MySQL
- Docker

## Run
1. Copy .env.example to .env
2. Set database credentials
2. Run `go run main.go`

## User Registration Feature
- Endpoint: POST /api/register
- Validation:
  - email format
  - password min length
  - duplicate email check
- Password hashed using bcrypt
- No JWT returned
