.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/proxy functions/payloadProxy.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupIndex functions/group/index.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupCreate functions/group/create.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupDestroy functions/group/destroy.go

	env GOOS=linux go build -ldflags="-s -w" -o main functions/payloadTransform.go
	zip bin/transform.zip main
	mv main bin/transform

	env GOOS=linux go build -ldflags="-s -w" -o main functions/group/index.go
	zip bin/group/index.zip main
	mv main bin/group/index

	env GOOS=linux go build -ldflags="-s -w" -o main functions/group/create.go
	zip bin/group/create.zip main
	mv main bin/group/create

	env GOOS=linux go build -ldflags="-s -w" -o main functions/group/destroy.go
	zip bin/group/destroy.zip main
	mv main bin/group/destroy

clean:
	rm -rf ./bin ./vendor go.sum

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
