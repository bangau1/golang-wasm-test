build:
	GOOS=js GOARCH=wasm go build -o docs/main.wasm

serve:
	goexec 'http.ListenAndServe(`:8081`, http.FileServer(http.Dir(`docs/`)))'