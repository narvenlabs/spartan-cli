build: generate build-test copy

build-test:
	goreleaser build --single-target --snapshot --rm-dist

build-prod:
	goreleaser build --rm-dist

generate:
	qtc

copy:
	yes | cp dist/igniter-cli_darwin_amd64/igniter /usr/local/bin/
