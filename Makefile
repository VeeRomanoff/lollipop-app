# Проверить наличие googleapis и клонирование при необходимости
GEN_DIR := .gen/proto/googleapis

# Правило для инициализации подмодуля
init-submodule:
	@echo "Initializing googleapis submodule..."
	@mkdir -p .gen/proto
	@if [ ! -d "$(GEN_DIR)/.git" ]; then \
		git submodule add --force https://github.com/googleapis/googleapis.git $(GEN_DIR); \
	fi
	@git submodule update --init --recursive
	@echo "Googleapis proto files ready"

$(GEN_DIR):
	@echo "Setting up googleapis submodule..."
	@mkdir -p .gen/proto
	@git submodule add https://github.com/googleapis/googleapis.git $(GEN_DIR) || true
	@git submodule update --init --recursive
	@echo "Googleapis proto files ready"

# Сгенерировать proto файлы
generate: init-submodule
	@echo "Generating files from proto..."
	@protoc -I. \
		-I$(GEN_DIR) \
		--go_out=./internal/pb/lollipop/ \
		--go-grpc_out=./internal/pb/lollipop/ \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		api/lollipop.proto \
		--grpc-gateway_out=./internal/pb/lollipop/ \
		--grpc-gateway_opt=paths=source_relative
	@echo "Completed √"

# Отчистка сгенерированных файлов
clean:
	@rm -rf internal/pb/lollipop/*.pb.go

distclean:
	@rm -rf .gen
	@git submodule deinit -f .gen/proto/googleapis || true
	@git rm -f .gen/proto/googleapis || true
	@echo "Cleaned √"

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


.PHONY: generate clean
