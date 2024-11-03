name := "cohnect"
commit := `git rev-parse --short HEAD`
build_time := `date -u +"%Y-%m-%dT%H:%M:%SZ"`
version := "0.0.1"

build:
	@echo "Building {{name}}..."
	go build -ldflags "-X main.Name={{name}} -X main.BuildTime={{build_time}}" -o {{name}}

run:
	@just build
	./{{name}}

version:
	@just build
	./{{name}} --version