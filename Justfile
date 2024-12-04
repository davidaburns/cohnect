name := "cohnect"
package_import := "github.com/davidaburns/cohnect"
commit := `git rev-parse --short HEAD`
version := "0.0.1"
build_time := `date -u +"%Y-%m-%dT%H:%M:%SZ"`

build type="dev":
	@just generate-buffers
	@echo "Building {{name}} ({{type}})..." ;
	@if [ "{{type}}" = "dev" ]; then \
		go build -ldflags "-X '{{package_import}}/config.name={{name}}' -X '{{package_import}}/config.buildTime={{build_time}}' -X '{{package_import}}/config.commit={{commit}}' -X '{{package_import}}/config.buildType={{type}}'" -o {{name}} ./cmd/{{name}}; \
	elif [ "{{type}}" = "prod" ]; then \
		go build -ldflags "-X '{{package_import}}/config.name={{name}}' -X '{{package_import}}/config.buildTime={{build_time}}' -X '{{package_import}}/config.commit={{commit}}' -X '{{package_import}}/config.buildType={{type}}' -s -w" -trimpath -o {{name}} ./cmd/{{name}}; \
	else \
		echo "Unknown build type {{type}}, please use 'dev' or 'prod'"; \
		exit 1; \
	fi

run type +args="":
	@just build {{type}}
	@./{{name}} {{args}}

generate-buffers:
	@echo "Generating flatbuffers..."
	@rm -rf ./internal/server/buffers
	@flatc --go --gen-object-api -o ./internal/server ./schemas/buffers.fbs;
	@find ./internal/server/buffers -type f -name '*.go' | while read file; do mv "$file" "$(dirname "$file")/$(basename "$file" | tr '[:upper:]' '[:lower:]')"; done

test:
	@echo "Running Tests..."
	@go test ./...

lint:
	@echo "Linting..."
	@golangci-lint run ./...

clean:
	@echo "Cleaning..."
	@rm ./{{name}}