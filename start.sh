go build -v -o dist/ ./cmd/...

gheye stubs --config stubs/config.yml

./dist/gheye stubs --config stubs/config.yml