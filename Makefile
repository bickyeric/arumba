build:
	go build -o output/updater app/update-collector/main.go
	go build -o output/webhook app/webhook/main.go

test:
	go test ./... -cover -count=1