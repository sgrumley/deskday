list:
		sqlite3 network_connections.db "SELECT * FROM connections;"

install:
		chmod +x ./scripts/setup.sh
		./scripts/setup.sh
		go build ./cmd/deskday/...

run:
	go run ./cmd/deskday/...
