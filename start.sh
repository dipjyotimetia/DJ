go build -v -o dist/ ./cmd/...

dj stubs --config stubs/config.yml

./dist/dj stubs --config stubs/config.yml

git tag -a v1.0.0 -m "First release" && git push origin v1.0.0