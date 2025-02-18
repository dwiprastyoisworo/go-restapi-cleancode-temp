.PHONY: run server migrate rollback build clean

# Menjalankan service utama
app:
	go run cmd/app.serve.go

# Menjalankan migrasi database (up)
migrate:
	go run cmd/migration.serve.go -type=run

# Menjalankan rollback migrasi (down)
rollback:
	go run cmd/migration.serve.go -type=rollback

# Build binary untuk service utama dan migrasi
build:
	go build -o bin/serve cmd/app.serve.go

# Membersihkan binary build
clean:
	rm -f bin/serve