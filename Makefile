.PHONY: dependency docker-up docker-down deploy clean docker-remove-image server-logs

dependency:
	@go get -v ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

docker-remove-image:
	@docker rmi creating-simple-api_web
	@docker volume rm creating-simple-api_my-db

logs:
	@docker-compose logs -f web

deploy: docker-up

clean: docker-down docker-remove-image
  