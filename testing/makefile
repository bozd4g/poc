run-server:
	@go run server/cmd/main.go

run-client:
	@go run client/cmd/main.go

consumer:
	@rm -rf client/internal/todo/pacts
	@go test ./... -tags=consumer -count=1

provider:
	@go test ./... -tags=provider -count=1

integration:
	@go test ./... -tags=integration -count=1

unit:
	@go test ./... -count=1

publish:
	@sh .pact/publish.sh