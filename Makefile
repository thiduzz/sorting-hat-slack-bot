.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupStore functions/group/store.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupIndex functions/group/index.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupSubscribe functions/group/subscribe.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
