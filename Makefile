.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/slack/slash/hats slack/slash/hats.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slack/interactivity/actions slack/interactivity/actions.go
clean:
	rm -rf ./bin ./vendor go.sum

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
