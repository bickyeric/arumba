BASE	:= github.com/bickyeric/arumba

build:
	go build -o output/updater app/update-collector/main.go
	go build -o output/webhook app/webhook/main.go

test:
	go test ./... -cover -count=1

generate:
	mockgen -destination=mocks/comic_repo.go -package=mocks ${BASE}/repository IComic
	mockgen -destination=mocks/episode_repo.go -package=mocks ${BASE}/repository IEpisode
	mockgen -destination=mocks/source_repo.go -package=mocks ${BASE}/repository ISource
