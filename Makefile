.PHONY: run server migrate rollback build clean

# Menjalankan service utama
run:
	go run cmd/app/*.go

# Menjalankan migrasi database (up)
migrate:
	go run cmd/migration/*.go -type=run

# Menjalankan rollback migrasi (down)
rollback:
	go run cmd/migration/*.go -type=rollback

# Build binary untuk service utama dan migrasi
build:
	go build -o bin/serve cmd/app/*.go

# Membersihkan binary build
clean:
	rm -f bin/serve