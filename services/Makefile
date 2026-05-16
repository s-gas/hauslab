up:
	docker image prune -f
	docker network create hauslab || true
	docker compose -f adguard/docker-compose.yaml up -d --pull always
	docker compose -f postgres/docker-compose.yaml up -d --pull always
	docker compose -f observability/docker-compose.yaml up -d --pull always
	docker compose -f sysmetrics/docker-compose.yaml up -d --pull always
	docker compose -f svcmonitor/docker-compose.yaml up -d --pull always
	docker compose -f gasmetrics/docker-compose.yaml up -d --pull always
	docker compose -f nginx/docker-compose.yaml up -d --pull always

down:
	docker compose -f nginx/docker-compose.yaml down
	docker compose -f gasmetrics/docker-compose.yaml down
	docker compose -f svcmonitor/docker-compose.yaml down
	docker compose -f sysmetrics/docker-compose.yaml down
	docker compose -f observability/docker-compose.yaml down
	docker compose -f postgres/docker-compose.yaml down
	docker compose -f adguard/docker-compose.yaml down
	docker network rm hauslab || true
