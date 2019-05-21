compile:
	go build -o output/updater app/update-collector/main.go
	go build -o output/webhook app/webhook/main.go