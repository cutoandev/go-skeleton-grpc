proto-gen-all:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/**/*.proto

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ${FILE}

install:
	go get
	make proto-gen-all

setup:
	make install
	docker-compose -f _docker/database-local.yml up -d 300592


# For database config
db-migration:
	atlas migrate diff --env gorm ${file}
	echo "[IMPORTANT] self-check generated file"

# For run code
run:
	make proto-gen-all
	go run main.go

build-local:
	docker rmi -f skeleton-grc-local:1.0
	docker-compose up --force-recreate
