exe:
	test -d bin || mkdir -p bin
	cd bin;go build ../proxy/main/proxy.go;cd ../
	cd bin;go build ../node/main/node.go; cd ../

proto:
	cd protocol/proto;protoc --go_out=. *.proto;mv *.go ../;cd ../../

proxy_start:
	 nohup bin/proxy configs/proxy/config.toml webhook > proxy.log 2>&1 &

node_start:
	 nohup bin/node configs/node/config.toml webhook > node.log 2>&1 &