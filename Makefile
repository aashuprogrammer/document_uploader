DB_URL=postgresql://neondb_owner:npg_9gUXz4LWjFmi@ep-tiny-glade-aih7e7m3-pooler.c-4.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require

dev:
	fiber dev

migrateup:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose down 1

migratedrop:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose drop -f

sqlc:
	sqlc generate

.PHONY: dev migrateup migrateup1 migratedown migratedown1 migratedrop sqlc