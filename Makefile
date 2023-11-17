dep:
	go mod download
dep-sync:
	go mod tidy

proto-generate:
	docker run --rm -v `pwd`/external:/defs namely/protoc-all:1.51_0 -i dto -f balances/Common.proto -l go -o ./
	docker run --rm -v `pwd`/external:/defs namely/protoc-all:1.51_0 -i dto -f balances/EmmitBalanceRequest.proto -l go -o ./