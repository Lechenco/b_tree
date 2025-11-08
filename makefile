build-app:
	go build -o build/app cmd/app/main.go 

start-app:
	./build/app

run:
	go run cmd/app/main.go