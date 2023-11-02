dev-client:
	BUILD_ENV=development npm --prefix ./views run dev

dev-server:
	BUILD_ENV=development go run main.go