up:
	docker network create sysmetrics-observability || true
	docker network create svcmonitor-observability || true
	docker network create monitor || true
	docker compose -f observability/docker-compose.yaml up -d
	docker compose -f sysmetrics/docker-compose.yaml up -d --build
	docker compose -f svcmonitor/docker-compose.yaml up -d --build

down:
	docker compose -f observability/docker-compose.yaml down
	docker compose -f svcmonitor/docker-compose.yaml down
	docker compose -f sysmetrics/docker-compose.yaml down
	docker network rm monitor || true
	docker network rm sysmetrics-observability || true
	docker network rm svcmonitor-observability || true
