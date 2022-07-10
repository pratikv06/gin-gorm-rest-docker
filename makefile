.PHONY: run install db-up

run:
	go run main.go

install:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres

db-up:
	docker-compose up --build