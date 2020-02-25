exe:
	test -d bin || mkdir -p bin
	cd bin;go build ../main/webhook.go;cd ../
	cd bin;go build ../client/main/git_client.go; cd ../

proto:
	cd codec/proto;protoc --go_out=. *.proto;mv *.go ../;cd ../../