run:
	docker compose up -d
	cd app && go run .

.PHONY: run
