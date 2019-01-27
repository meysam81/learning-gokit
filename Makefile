run-server:
	go run `ls | grep ".go"` &

run-client:
	curl localhost -d '{"s": "hello meysam"}' localhost:8080/uppercase 2>/dev/null
	curl localhost -d '{"s": "hello meysam"}' localhost:8080/count 2>/dev/null
