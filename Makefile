default: out/example

clean:
	rm -rf out

test: *.go
	go test

out/example: implementation.go cmd/example/main.go
	go get github.com/ProMKQ/kpi-lab2
	mkdir -p out
	go build -o out/example ./cmd/example
