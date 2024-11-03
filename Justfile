name := "cohnect"
commit := `git rev-parse --short HEAD`
version := "0.0.1"
build_time := `date -u +"%Y-%m-%dT%H:%M:%SZ"`
build_path := "./build"

build type="dev":
	@echo "Building {{name}} ({{type}})..." ;
	@if [ "{{type}}" = "dev" ]; then \
		go build -ldflags "-X 'main.Name={{name}}' -X 'main.BuildTime={{build_time}}' -X 'main.Commit={{commit}}' -X 'main.BuildType={{type}}'" -o {{build_path}}/{{type}}/{{name}}; \
	elif [ "{{type}}" = "prod" ]; then \
		go build -ldflags "-X 'main.Name={{name}}' -X 'main.BuildTime={{build_time}}' -X 'main.Commit={{commit}}' -X 'main.BuildType={{type}}' -s -w" -trimpath -o {{build_path}}/{{type}}/{{name}}; \
	else \
		echo "Unknown build type {{type}}, please use 'dev' or 'prod'"; \
		exit 1; \
	fi

run:
	@just build
	./{{name}}

version type="dev":
	@just build {{type}}
	@{{build_path}}/{{type}}/{{name}} --version

test:
	@echo "Running Tests..."
	@go test ./...

lint:
	@echo "Linting..."
	@golangci-lint run ./...

clean:
	@echo "Cleaning..."
	@rm -rf {{build_path}}