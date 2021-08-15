// Code generated by qtc from "project.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/project.qtpl:1
package templates

//line templates/project.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/project.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/project.qtpl:1
func StreamGenREADME(qw422016 *qt422016.Writer, name string) {
//line templates/project.qtpl:1
	qw422016.N().S(`#  `)
//line templates/project.qtpl:2
	qw422016.E().S(name)
//line templates/project.qtpl:2
	qw422016.N().S(`

> This project was generated using [igniter-cli](https://github.com/Narven/igniter-cli)

`)
//line templates/project.qtpl:6
}

//line templates/project.qtpl:6
func WriteGenREADME(qq422016 qtio422016.Writer, name string) {
//line templates/project.qtpl:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:6
	StreamGenREADME(qw422016, name)
//line templates/project.qtpl:6
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:6
}

//line templates/project.qtpl:6
func GenREADME(name string) string {
//line templates/project.qtpl:6
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:6
	WriteGenREADME(qb422016, name)
//line templates/project.qtpl:6
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:6
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:6
	return qs422016
//line templates/project.qtpl:6
}

//line templates/project.qtpl:8
func StreamGenGitignore(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:8
	qw422016.N().S(`dist
build
tmp
.env
*.log

# Editors
.idea
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
*.sublime*

# System Files
.DS_Store
Thumbs.db

node_modules

# Go
vendor/
`)
//line templates/project.qtpl:32
}

//line templates/project.qtpl:32
func WriteGenGitignore(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:32
	StreamGenGitignore(qw422016)
//line templates/project.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:32
}

//line templates/project.qtpl:32
func GenGitignore() string {
//line templates/project.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:32
	WriteGenGitignore(qb422016)
//line templates/project.qtpl:32
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:32
	return qs422016
//line templates/project.qtpl:32
}

//line templates/project.qtpl:34
func StreamGenGoMod(qw422016 *qt422016.Writer, moduleName string) {
//line templates/project.qtpl:34
	qw422016.N().S(`module `)
//line templates/project.qtpl:35
	qw422016.E().S(moduleName)
//line templates/project.qtpl:35
	qw422016.N().S(`

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/spf13/viper v1.8.1
)

`)
//line templates/project.qtpl:46
}

//line templates/project.qtpl:46
func WriteGenGoMod(qq422016 qtio422016.Writer, moduleName string) {
//line templates/project.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:46
	StreamGenGoMod(qw422016, moduleName)
//line templates/project.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:46
}

//line templates/project.qtpl:46
func GenGoMod(moduleName string) string {
//line templates/project.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:46
	WriteGenGoMod(qb422016, moduleName)
//line templates/project.qtpl:46
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:46
	return qs422016
//line templates/project.qtpl:46
}

//line templates/project.qtpl:48
func StreamGenDockerCompose(qw422016 *qt422016.Writer, projectName, dbDriver string) {
//line templates/project.qtpl:48
	qw422016.N().S(`version: '3.8'
services:
  database:
    container_name: `)
//line templates/project.qtpl:52
	qw422016.E().S(projectName)
//line templates/project.qtpl:52
	qw422016.N().S(`-database
    env_file:
      - .env
`)
//line templates/project.qtpl:55
	if dbDriver == "mysql" {
//line templates/project.qtpl:55
		qw422016.N().S(`    image: mysql:8
    command: --default-authentication-plugin=mysql_native_password --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
    volumes:
      - data:/var/lib/mysql
    restart: always
    environment:
      MYSQL_DATABASE: `)
//line templates/project.qtpl:62
		qw422016.E().S(projectName)
//line templates/project.qtpl:62
		qw422016.N().S(`_dev
      MYSQL_USER: `)
//line templates/project.qtpl:63
		qw422016.E().S(projectName)
//line templates/project.qtpl:63
		qw422016.N().S(`
      MYSQL_PASSWORD: `)
//line templates/project.qtpl:64
		qw422016.E().S(projectName)
//line templates/project.qtpl:64
		qw422016.N().S(`
      MYSQL_ROOT_PASSWORD: root
    ports:
      - 3306:3306
    networks:
      - `)
//line templates/project.qtpl:69
		qw422016.E().S(projectName)
//line templates/project.qtpl:69
		qw422016.N().S(`_network
`)
//line templates/project.qtpl:70
	} else if dbDriver == "postgres" {
//line templates/project.qtpl:70
		qw422016.N().S(`    image: postgres
    ports:
      - "5432:5432"
    volumes:
      - data:/var/lib/postgresql/data
`)
//line templates/project.qtpl:76
	}
//line templates/project.qtpl:76
	qw422016.N().S(`
volumes:
  data:

networks:
  `)
//line templates/project.qtpl:82
	qw422016.E().S(projectName)
//line templates/project.qtpl:82
	qw422016.N().S(`_network:
    driver: bridge
`)
//line templates/project.qtpl:84
}

//line templates/project.qtpl:84
func WriteGenDockerCompose(qq422016 qtio422016.Writer, projectName, dbDriver string) {
//line templates/project.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:84
	StreamGenDockerCompose(qw422016, projectName, dbDriver)
//line templates/project.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:84
}

//line templates/project.qtpl:84
func GenDockerCompose(projectName, dbDriver string) string {
//line templates/project.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:84
	WriteGenDockerCompose(qb422016, projectName, dbDriver)
//line templates/project.qtpl:84
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:84
	return qs422016
//line templates/project.qtpl:84
}

//line templates/project.qtpl:86
func StreamGenEnv(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:86
	qw422016.N().S(`DASHBOARD_URL=
DASHBOARD_HOST=
DATABASE_NAME=
DATABASE_PASSWORD=
DATABASE_PORT=
DATABASE_USER=
`)
//line templates/project.qtpl:93
}

//line templates/project.qtpl:93
func WriteGenEnv(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:93
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:93
	StreamGenEnv(qw422016)
//line templates/project.qtpl:93
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:93
}

//line templates/project.qtpl:93
func GenEnv() string {
//line templates/project.qtpl:93
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:93
	WriteGenEnv(qb422016)
//line templates/project.qtpl:93
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:93
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:93
	return qs422016
//line templates/project.qtpl:93
}

//line templates/project.qtpl:95
func StreamGenMakeFile(qw422016 *qt422016.Writer, name string) {
//line templates/project.qtpl:95
	qw422016.N().S(`.RECIPEPREFIX = >
.PHONY: all
.DEFAULT_GOAL := help

BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

# This how we want to name the binary output
BINARY=`)
//line templates/project.qtpl:107
	qw422016.E().S(name)
//line templates/project.qtpl:107
	qw422016.N().S(`

# These are the values we want to pass for VERSION and BUILD
VERSION=1.0.0
BUILD=`)
//line templates/project.qtpl:107
	qw422016.N().S("`")
//line templates/project.qtpl:107
	qw422016.N().S(`git rev-parse HEAD`)
//line templates/project.qtpl:107
	qw422016.N().S("`")
//line templates/project.qtpl:107
	qw422016.N().S(`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

start:
	> air -c .air.toml
	# > go run api/main.go

build:
	> go build ${LDFLAGS} -o ${BINARY}

deps: ## Download dependencies
	> go mod tidy

fmt: ## Format source code with gofmt
	> find . -name "*.go" -exec gofmt -s -w {} \;

help:
	> @echo "Help: `)
//line templates/project.qtpl:130
	qw422016.E().S(name)
//line templates/project.qtpl:130
	qw422016.N().S(` root Makefile"
	> @echo "Usage: make [TARGET] [EXTRA_ARGUMENTS]"
	> @echo "Targets:"
	> @echo "~> start					- starts the application"

`)
//line templates/project.qtpl:135
}

//line templates/project.qtpl:135
func WriteGenMakeFile(qq422016 qtio422016.Writer, name string) {
//line templates/project.qtpl:135
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:135
	StreamGenMakeFile(qw422016, name)
//line templates/project.qtpl:135
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:135
}

//line templates/project.qtpl:135
func GenMakeFile(name string) string {
//line templates/project.qtpl:135
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:135
	WriteGenMakeFile(qb422016, name)
//line templates/project.qtpl:135
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:135
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:135
	return qs422016
//line templates/project.qtpl:135
}

//line templates/project.qtpl:137
func StreamGenEditorConfig(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:137
	qw422016.N().S(`root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true
indent_style = space
indent_size = 4

[Makefile]
indent_style = tab

[*.go]
indent_style = tab

[*.css]
indent_size = 2

[*.toml]
indent_size = 2

[*.js]
indent_size = 2
block_comment_start = /*
block_comment_end = */

[*.{html,htm}]
indent_size = 2

[*.{yml,yaml}]
indent_size = 2

[*.json]
indent_size = 2

[*.diff]
indent_size = 1

`)
//line templates/project.qtpl:177
}

//line templates/project.qtpl:177
func WriteGenEditorConfig(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:177
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:177
	StreamGenEditorConfig(qw422016)
//line templates/project.qtpl:177
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:177
}

//line templates/project.qtpl:177
func GenEditorConfig() string {
//line templates/project.qtpl:177
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:177
	WriteGenEditorConfig(qb422016)
//line templates/project.qtpl:177
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:177
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:177
	return qs422016
//line templates/project.qtpl:177
}

//line templates/project.qtpl:179
func StreamGenDockerIgnore(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:179
	qw422016.N().S(`*.log
bin/
!go.sum
vendor/
*.test
coverage.out
coverage.txt
`)
//line templates/project.qtpl:187
}

//line templates/project.qtpl:187
func WriteGenDockerIgnore(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:187
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:187
	StreamGenDockerIgnore(qw422016)
//line templates/project.qtpl:187
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:187
}

//line templates/project.qtpl:187
func GenDockerIgnore() string {
//line templates/project.qtpl:187
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:187
	WriteGenDockerIgnore(qb422016)
//line templates/project.qtpl:187
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:187
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:187
	return qs422016
//line templates/project.qtpl:187
}

//line templates/project.qtpl:189
func StreamGenDockerfile(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:189
	qw422016.N().S(`FROM golang:1.16.2-alpine3.13 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.16.2-alpine3.13 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]
`)
//line templates/project.qtpl:208
}

//line templates/project.qtpl:208
func WriteGenDockerfile(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:208
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:208
	StreamGenDockerfile(qw422016)
//line templates/project.qtpl:208
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:208
}

//line templates/project.qtpl:208
func GenDockerfile() string {
//line templates/project.qtpl:208
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:208
	WriteGenDockerfile(qb422016)
//line templates/project.qtpl:208
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:208
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:208
	return qs422016
//line templates/project.qtpl:208
}

//line templates/project.qtpl:210
func StreamGenerateApiMain(qw422016 *qt422016.Writer, moduleName, projectName, dbDriver string) {
//line templates/project.qtpl:210
	qw422016.N().S(`package main

import (
    "`)
//line templates/project.qtpl:214
	qw422016.E().S(moduleName)
//line templates/project.qtpl:214
	qw422016.N().S(`/config"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/jmoiron/sqlx"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("LoadConfig error: %v", err)
    }

    `)
//line templates/project.qtpl:232
	if dbDriver == "mysql" {
//line templates/project.qtpl:232
		qw422016.N().S(`
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?parseTime=true",
        cfg.DatabaseUser,
        cfg.DatabasePassword,
        cfg.DatabaseHost,
        cfg.DatabasePort,
        cfg.DatabaseName,
    )
    `)
//line templates/project.qtpl:241
	} else if dbDriver == "postgres" {
//line templates/project.qtpl:241
		qw422016.N().S(`

    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?sslmode=disable",
        cfg.DatabaseUser,
        cfg.DatabasePassword,
        cfg.DatabaseHost,
        cfg.DatabasePort,
        cfg.DatabaseName,
    )
    `)
//line templates/project.qtpl:251
	}
//line templates/project.qtpl:251
	qw422016.N().S(`

    db, err := sqlx.Connect("`)
//line templates/project.qtpl:253
	qw422016.E().S(dbDriver)
//line templates/project.qtpl:253
	qw422016.N().S(`", dsn)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer db.Close()

    app := fiber.New(fiber.Config{
        CaseSensitive: true,
    })
    app.Use(fiberLogger.New())
    app.Use(recover.New())

    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
        AllowHeaders:     "Authorization,Content-Type,Crossdomain,Origin",
        AllowMethods:     "PUT,PATCH,GET,POST,DELETE",
        AllowOrigins:     "*",
        ExposeHeaders:    "Content-Length,Authorization",
    }))

    app.Get("/ping", func(ctx *fiber.Ctx) error {
        return ctx.SendStatus(fiber.StatusOK)
    })

    log.Fatal(app.Listen(fmt.Sprintf(":%d", 9069)))
}
`)
//line templates/project.qtpl:279
}

//line templates/project.qtpl:279
func WriteGenerateApiMain(qq422016 qtio422016.Writer, moduleName, projectName, dbDriver string) {
//line templates/project.qtpl:279
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:279
	StreamGenerateApiMain(qw422016, moduleName, projectName, dbDriver)
//line templates/project.qtpl:279
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:279
}

//line templates/project.qtpl:279
func GenerateApiMain(moduleName, projectName, dbDriver string) string {
//line templates/project.qtpl:279
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:279
	WriteGenerateApiMain(qb422016, moduleName, projectName, dbDriver)
//line templates/project.qtpl:279
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:279
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:279
	return qs422016
//line templates/project.qtpl:279
}

//line templates/project.qtpl:281
func StreamGenWelcomeMsg(qw422016 *qt422016.Writer, name string) {
//line templates/project.qtpl:281
	qw422016.N().S(`
🔥️ Project `)
//line templates/project.qtpl:282
	qw422016.E().S(name)
//line templates/project.qtpl:282
	qw422016.N().S(` generated successfully

Next steps:

    * cd `)
//line templates/project.qtpl:286
	qw422016.E().S(name)
//line templates/project.qtpl:286
	qw422016.N().S(`
    * make start
`)
//line templates/project.qtpl:288
}

//line templates/project.qtpl:288
func WriteGenWelcomeMsg(qq422016 qtio422016.Writer, name string) {
//line templates/project.qtpl:288
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:288
	StreamGenWelcomeMsg(qw422016, name)
//line templates/project.qtpl:288
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:288
}

//line templates/project.qtpl:288
func GenWelcomeMsg(name string) string {
//line templates/project.qtpl:288
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:288
	WriteGenWelcomeMsg(qb422016, name)
//line templates/project.qtpl:288
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:288
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:288
	return qs422016
//line templates/project.qtpl:288
}

//line templates/project.qtpl:290
func StreamGenBaseEntity(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:290
	qw422016.N().S(`package entity

type ID = int64
`)
//line templates/project.qtpl:294
}

//line templates/project.qtpl:294
func WriteGenBaseEntity(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:294
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:294
	StreamGenBaseEntity(qw422016)
//line templates/project.qtpl:294
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:294
}

//line templates/project.qtpl:294
func GenBaseEntity() string {
//line templates/project.qtpl:294
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:294
	WriteGenBaseEntity(qb422016)
//line templates/project.qtpl:294
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:294
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:294
	return qs422016
//line templates/project.qtpl:294
}

//line templates/project.qtpl:296
func StreamGenMITLicense(qw422016 *qt422016.Writer, year int) {
//line templates/project.qtpl:296
	qw422016.N().S(`MIT License

Copyright (c) `)
//line templates/project.qtpl:299
	qw422016.N().D(year)
//line templates/project.qtpl:299
	qw422016.N().S(`

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`)
//line templates/project.qtpl:318
}

//line templates/project.qtpl:318
func WriteGenMITLicense(qq422016 qtio422016.Writer, year int) {
//line templates/project.qtpl:318
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:318
	StreamGenMITLicense(qw422016, year)
//line templates/project.qtpl:318
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:318
}

//line templates/project.qtpl:318
func GenMITLicense(year int) string {
//line templates/project.qtpl:318
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:318
	WriteGenMITLicense(qb422016, year)
//line templates/project.qtpl:318
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:318
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:318
	return qs422016
//line templates/project.qtpl:318
}

//line templates/project.qtpl:320
func StreamGenGithubIssueTemplate(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:320
	qw422016.N().S(`### Issue report

`)
//line templates/project.qtpl:323
}

//line templates/project.qtpl:323
func WriteGenGithubIssueTemplate(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:323
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:323
	StreamGenGithubIssueTemplate(qw422016)
//line templates/project.qtpl:323
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:323
}

//line templates/project.qtpl:323
func GenGithubIssueTemplate() string {
//line templates/project.qtpl:323
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:323
	WriteGenGithubIssueTemplate(qb422016)
//line templates/project.qtpl:323
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:323
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:323
	return qs422016
//line templates/project.qtpl:323
}

//line templates/project.qtpl:325
func StreamGenPullRequestTemplate(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:325
	qw422016.N().S(`### What did you implement:

Closes #XXXXX

### How did you implement it:

...

### How can we verify it:

...

### TODO's:

- [ ] Write documentation
- [ ] Check that there aren't other open pull requests for the same issue/feature
- [ ] Format your source code by `)
//line templates/project.qtpl:325
	qw422016.N().S("`")
//line templates/project.qtpl:325
	qw422016.N().S(`make fmt`)
//line templates/project.qtpl:325
	qw422016.N().S("`")
//line templates/project.qtpl:325
	qw422016.N().S(`
- [ ] Pass the test by `)
//line templates/project.qtpl:325
	qw422016.N().S("`")
//line templates/project.qtpl:325
	qw422016.N().S(`make test`)
//line templates/project.qtpl:325
	qw422016.N().S("`")
//line templates/project.qtpl:325
	qw422016.N().S(`
- [ ] Provide verification config / commands
- [ ] Enable "Allow edits from maintainers" for this PR
- [ ] Update the messages below

**Is this ready for review?:** No

**Is it a breaking change?:** No
`)
//line templates/project.qtpl:351
}

//line templates/project.qtpl:351
func WriteGenPullRequestTemplate(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:351
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:351
	StreamGenPullRequestTemplate(qw422016)
//line templates/project.qtpl:351
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:351
}

//line templates/project.qtpl:351
func GenPullRequestTemplate() string {
//line templates/project.qtpl:351
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:351
	WriteGenPullRequestTemplate(qb422016)
//line templates/project.qtpl:351
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:351
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:351
	return qs422016
//line templates/project.qtpl:351
}

//line templates/project.qtpl:353
func StreamGenGithubCI(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:353
	qw422016.N().S(`name: CI

on:
  push:
    branches: [master, develop]
  pull_request:
    branches: [master]

jobs:
  Go:
    name: Go
    runs-on: ubuntu-latest

    env:
      SRC_DIR: src/github.com/${{ github.repository }}

    strategy:
      matrix:
        go: [ '1.13.x', '1.14.x', '1.15.x' ]

    steps:
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: ${{ matrix.go }}
        id: go

      - name: Setup PATH
        run: |
          echo "GOPATH=${{ github.workspace }}" >> "$GITHUB_ENV"
          echo "GOBIN=${{ github.workspace }}/bin" >> "$GITHUB_ENV"
          echo "${{ github.workspace }}/bin" >> "$GITHUB_PATH"

      - name: Checkout
        uses: actions/checkout@v2
        with:
          path: ${{env.SRC_DIR}}

      - name: Download dependencies
        working-directory: ${{env.SRC_DIR}}
        run: make deps

      - name: Build binary
        working-directory: ${{env.SRC_DIR}}
        run: make build

      - name: Run tests
        working-directory: ${{env.SRC_DIR}}
        run: make tests

`)
//line templates/project.qtpl:404
}

//line templates/project.qtpl:404
func WriteGenGithubCI(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:404
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:404
	StreamGenGithubCI(qw422016)
//line templates/project.qtpl:404
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:404
}

//line templates/project.qtpl:404
func GenGithubCI() string {
//line templates/project.qtpl:404
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:404
	WriteGenGithubCI(qb422016)
//line templates/project.qtpl:404
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:404
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:404
	return qs422016
//line templates/project.qtpl:404
}

//line templates/project.qtpl:406
func StreamGenTravis(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:406
	qw422016.N().S(`language: go
sudo: false

matrix:
  include:
    - go: 1.x
      env: LATEST=true
    - go: 1.15.x
    - go: tip
  allow_failures:
    - go: tip

install:
  - export GO111MODULE=on
  - go get -t -v $(go list ./... | grep -v -E "vendor")

script:
  - go test -cover -coverprofile=coverage.txt -covermode=atomic -v $(go list ./... | grep -v -E "vendor")

# after_success:
#  - bash <(curl -s https://codecov.io/bash)

# notifications:
#  email: false
`)
//line templates/project.qtpl:431
}

//line templates/project.qtpl:431
func WriteGenTravis(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:431
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:431
	StreamGenTravis(qw422016)
//line templates/project.qtpl:431
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:431
}

//line templates/project.qtpl:431
func GenTravis() string {
//line templates/project.qtpl:431
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:431
	WriteGenTravis(qb422016)
//line templates/project.qtpl:431
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:431
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:431
	return qs422016
//line templates/project.qtpl:431
}

//line templates/project.qtpl:433
func StreamGenAirToml(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:433
	qw422016.N().S(`# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "."
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use `)
//line templates/project.qtpl:433
	qw422016.N().S("`")
//line templates/project.qtpl:433
	qw422016.N().S(`make`)
//line templates/project.qtpl:433
	qw422016.N().S("`")
//line templates/project.qtpl:433
	qw422016.N().S(` as well.
cmd = "go build -o ./tmp/main ."
# Binary file yields from `)
//line templates/project.qtpl:433
	qw422016.N().S("`")
//line templates/project.qtpl:433
	qw422016.N().S(`cmd`)
//line templates/project.qtpl:433
	qw422016.N().S("`")
//line templates/project.qtpl:433
	qw422016.N().S(`.
bin = "tmp/main"
# Customize binary.
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# Watch these filename extensions.
include_ext = ["go", "html"]
# Ignore these filename extensions or directories.
exclude_dir = ["tmp", "vendor", ".github", "migrations", "bin", "dist"]
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = [".env", "README.md", ".editconfig", "Makefile"]
# Exclude specific regular expressions.
exclude_regex = ["_test.go"]
# Exclude unchanged files.
exclude_unchanged = true
# Follow symlink for directories
follow_symlink = true
# This log file places in your tmp_dir.
log = "air.log"
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 1000 # ms
# Stop running old binary when build errors occur.
stop_on_error = true
# Send Interrupt signal before killing process (windows does not support this feature)
send_interrupt = false
# Delay after sending Interrupt signal
kill_delay = 500 # ms

[log]
# Show log time
time = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true
`)
//line templates/project.qtpl:485
}

//line templates/project.qtpl:485
func WriteGenAirToml(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:485
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:485
	StreamGenAirToml(qw422016)
//line templates/project.qtpl:485
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:485
}

//line templates/project.qtpl:485
func GenAirToml() string {
//line templates/project.qtpl:485
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:485
	WriteGenAirToml(qb422016)
//line templates/project.qtpl:485
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:485
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:485
	return qs422016
//line templates/project.qtpl:485
}

//line templates/project.qtpl:487
func StreamGenAuthors(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:487
	qw422016.N().S(`# This file lists all individuals having contributed content to the repository.

FirstName LastName <FirstName.LastName@provider.com>
`)
//line templates/project.qtpl:491
}

//line templates/project.qtpl:491
func WriteGenAuthors(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:491
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:491
	StreamGenAuthors(qw422016)
//line templates/project.qtpl:491
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:491
}

//line templates/project.qtpl:491
func GenAuthors() string {
//line templates/project.qtpl:491
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:491
	WriteGenAuthors(qb422016)
//line templates/project.qtpl:491
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:491
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:491
	return qs422016
//line templates/project.qtpl:491
}

//line templates/project.qtpl:493
func StreamGenGitAttributes(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:493
	qw422016.N().S(`# Auto detect text files and perform LF normalization
* text=auto

# Reduce conflicts on markdown files
*.md merge=union
`)
//line templates/project.qtpl:499
}

//line templates/project.qtpl:499
func WriteGenGitAttributes(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:499
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:499
	StreamGenGitAttributes(qw422016)
//line templates/project.qtpl:499
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:499
}

//line templates/project.qtpl:499
func GenGitAttributes() string {
//line templates/project.qtpl:499
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:499
	WriteGenGitAttributes(qb422016)
//line templates/project.qtpl:499
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:499
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:499
	return qs422016
//line templates/project.qtpl:499
}

//line templates/project.qtpl:501
func StreamGenProjectConfig(qw422016 *qt422016.Writer) {
//line templates/project.qtpl:501
	qw422016.N().S(`package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseHost           string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"DATABASE_HOST" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	DatabaseName           string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"DATABASE_NAME" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	DatabasePassword       string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"DATABASE_PASSWORD" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	DatabasePort           string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"DATABASE_PORT" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	DatabaseUser           string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"DATABASE_USER" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	JWTKeyPath             string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"JWT_KEY_PATH" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	SendgridAPIKey         string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"SENDGRID_API_KEY" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	SendgridEmail          string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"SENDGRID_EMAIL" validate:"required"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	SentryDSN              string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"SENTRY_DSN"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	SentryRelease          string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"SENTRY_RELEASE"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
	SentryTracesSampleRate string `)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`mapstructure:"SENTRY_TRACES_SAMPLE_RATE"`)
//line templates/project.qtpl:501
	qw422016.N().S("`")
//line templates/project.qtpl:501
	qw422016.N().S(`
}

func LoadConfig() (config Config, err error) {
	v := viper.New()
	v.SetConfigType("env")
	v.SetConfigFile(".env")
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		fmt.Printf("Error when Fetching Configuration - %s", err)
	}

	if err := v.Unmarshal(&config); err != nil {
		return config, err
	}
	return
}
`)
//line templates/project.qtpl:540
}

//line templates/project.qtpl:540
func WriteGenProjectConfig(qq422016 qtio422016.Writer) {
//line templates/project.qtpl:540
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/project.qtpl:540
	StreamGenProjectConfig(qw422016)
//line templates/project.qtpl:540
	qt422016.ReleaseWriter(qw422016)
//line templates/project.qtpl:540
}

//line templates/project.qtpl:540
func GenProjectConfig() string {
//line templates/project.qtpl:540
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/project.qtpl:540
	WriteGenProjectConfig(qb422016)
//line templates/project.qtpl:540
	qs422016 := string(qb422016.B)
//line templates/project.qtpl:540
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/project.qtpl:540
	return qs422016
//line templates/project.qtpl:540
}
