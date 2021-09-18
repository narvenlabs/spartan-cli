.PHONY: all

build: generate build-test copy

.DEFAULT_GOAL := help

build-test:
	goreleaser build --single-target --snapshot --rm-dist

build-prod:
	goreleaser build --rm-dist

generate:
	qtc

copy:
	yes | cp dist/spartan-cli_darwin_amd64/spartan /usr/local/bin/

help:
	@echo "Help: Spartan CLI root Makefile"
	@echo "Usage: make [TARGET] [EXTRA_ARGUMENTS]"
	@echo "Targets:"
	@echo "~> build          - uses gorealeaser to build a test version and copies to /usr/local/bin"
	@echo "~> build-prod     - builds production version"
	@echo "~> templates            - generates templates, using quicktemplate"