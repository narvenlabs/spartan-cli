build: clean build-test

clean:
	rm -rf ./dist

build-test:
	goreleaser build --single-target --snapshot

build-prod:
	goreleaser build