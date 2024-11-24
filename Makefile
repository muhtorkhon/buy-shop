CURRENT_DIR=$(shell pwd)

-include .env
run:
		go run main.go

# PostgreSQL uchun DB URL
DB_URL="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DB)?sslmode=disable"

# DB_URL="postgres://developer:19966@localhost:5432/i_shop_db?sslmode=disable"

# Migratsiyalarni bajarish (up)
migrate_up:
		migrate -path migrations -database "$(DB_URL)" -verbose up

# Migratsiyalarni teskari (down)
migrate-down:
		migrate -path migrations -database "$(DB_URL)" -verbose down

# Migratsiyalarni barcha (up/down) bajarish
migrate-reset:
		migrate -path migrations/ -database "$(DB_URL)" down
		migrate -path migrations/ -database "$(DB_URL)" up

# Migratsiya holatini ko'rish
migrate-status:
		migrate -path migrations/ -database "$(DB_URL)" status

swaggo:
	swag init