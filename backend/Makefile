docker_prepare: 
	docker pull postgres:15.2-alpine
	docker network create cinema-network

postgres:
	docker run --name postgres15cinema --network cinema-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=cinema -p 2345:5432 -d postgres:15.2-alpine

server: 
	go run main.go

docker_clean: 
	docker stop postgres15cinema
	docker rm postgres15cinema
	docker rmi postgres:15.2-alpine
	docker stop seatmapbackend
	docker rm seatmapbackend
	docker rmi seatmapbackend
	docker network rm seatmap-network

.PHONY: docker_prepare postgres server docker_clean