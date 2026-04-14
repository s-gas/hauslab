build:
	cd sysmetrics && go build -o sysmetrics .
	cd svcmonitor && go build -o svcmonitor .
	docker compose -f observability/docker-compose.yaml build

up:
	./sysmetrics/sysmetrics &
	./svcmonitor/svcmonitor &
	docker compose -f observability/docker-compose.yaml up -d

down:
	pkill sysmetrics
	pkill svcmonitor
	docker compose -f observability/docker-compose.yaml down
