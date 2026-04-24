up:
	docker network create monitor || true
	docker compose -f observability/docker-compose.yaml up -d
	docker compose -f svcmonitor/docker-compose.yaml up -d --build
	docker compose -f sysmetrics/docker-compose.yaml up -d --build

down:
	docker compose -f observability/docker-compose.yaml down
	docker compose -f svcmonitor/docker-compose.yaml down
	docker compose -f sysmetrics/docker-compose.yaml down
	docker network rm monitor || true
