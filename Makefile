.PHONY: clean

clean:
	rm -rf dist/

build:
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js dist/
	GOOS=js GOARCH=wasm go build -o dist/main.wasm

update-wasm:
	GOOS=js GOARCH=wasm go build -o dist/main.wasm