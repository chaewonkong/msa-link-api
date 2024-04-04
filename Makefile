ref:
	@echo "Tidy and vendor"
	@go mod tidy && go mod vendor
PHONY: ref

run:
	@echo "Running the application"
	 @export $$(cat .env | xargs)
	@go run main.go
PHONY: run