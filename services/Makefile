up:
	docker network create monitor || true
	docker network create sysmetrics-observability || true
	docker compose -f observability/docker-compose.yaml up -d --pull always
	docker compose -f svcmonitor/docker-compose.yaml up -d --build --force-recreate
	docker compose -f sysmetrics/docker-compose.yaml up -d --build --force-recreate

down:
	docker compose -f observability/docker-compose.yaml down
	docker compose -f svcmonitor/docker-compose.yaml down
	docker compose -f sysmetrics/docker-compose.yaml down
	docker network rm monitor || true
	docker network rm sysmetrics-observability || true
