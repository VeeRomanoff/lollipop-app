# Сгенерировать proto файлы
generate:
	@echo "generating files from proto..."
	@protoc -I. -I./vendor/protogen --go_out=./internal/pb/lollipop \
	--go-grpc_out=./internal/pb/lollipop --go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative api/lollipop.proto \
	--grpc-gateway_out ./internal/pb/lollipop --grpc-gateway_opt paths=source_relative \

	@echo "completed √"

# Отчистка сгенерированных файлов
clean:
	@rm -rf internal/pb/lollipop/*.pb.go

# Установка зависимостей
install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

vendor:
	go mod vendor
	cp -a /Users/jamesharley/Desktop/googleapis ./vendor/protogen

# Отчиска сгенерированных файлов и генерация кода заново
rebuild: clean generate

run:
	go build cmd/main.go
	go run cmd/main.go

