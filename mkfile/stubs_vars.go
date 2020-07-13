package mkfile

const varsStub = `SHELL = bash

APP := $(shell basename $(PWD) | tr '[:upper:]' '[:lower:]')
TTY := $(shell if [ -z $$CI ]; then echo "-t"; fi)

GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)
GIT_SHORT_COMMIT_HASH := $(shell git rev-parse --short HEAD)`
