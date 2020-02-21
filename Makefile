build:
	GOOS=js GOARCH=wasm go build -o main.wasm

serve:
	goexec 'http.ListenAndServe(`:8081`, http.FileServer(http.Dir(`.`)))'