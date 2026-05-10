up:
	docker network create hauslab || true
	docker compose -f adguard/docker-compose.yaml up -d --pull always
	docker compose -f observability/docker-compose.yaml up -d --pull always
	docker compose -f sysmetrics/docker-compose.yaml up -d --pull always
	docker compose -f svcmonitor/docker-compose.yaml up -d --pull always
	docker compose -f nginx/docker-compose.yaml up -d --pull always

down:
	docker compose -f nginx/docker-compose.yaml down
	docker compose -f adguard/docker-compose.yaml down
	docker compose -f observability/docker-compose.yaml down
	docker compose -f svcmonitor/docker-compose.yaml down
	docker compose -f sysmetrics/docker-compose.yaml down
	docker network rm hauslab || true
