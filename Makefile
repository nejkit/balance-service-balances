dep:
	go mod download
dep-sync:
	go mod tidy

proto-generate:
	docker run --rm -v `pwd`/external:/defs namely/protoc-all:1.51_0 -d dto/balances -l go -o ./