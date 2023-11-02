dev-client:
	BUILD_ENV=development npm --prefix ./views run dev

dev-server:
	BUILD_ENV=development go run main.go

dev:
	(trap 'kill 0' SIGINT; BUILD_ENV=development go run main.go & BUILD_ENV=development npm --prefix ./views run dev & wait)