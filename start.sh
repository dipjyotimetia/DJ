go build -v -o dist/ ./cmd/...

dj stubs --config stubs/config.yml

./dist/dj stubs --config stubs/config.yml