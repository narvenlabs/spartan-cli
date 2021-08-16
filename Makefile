build: clean generate build-test copy

clean:
	rm -rf ./dist

build-test:
	goreleaser build --single-target --snapshot

build-prod:
	goreleaser build

generate:
	qtc

copy:
	yes | cp dist/igniter-cli_darwin_amd64/igniter /usr/local/bin/
