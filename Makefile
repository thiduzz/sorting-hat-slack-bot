.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/proxy functions/payloadProxy.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupIndex functions/group/index.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupCreate functions/group/create.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupDestroy functions/group/destroy.go
clean:
	rm -rf ./bin ./vendor go.sum

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
