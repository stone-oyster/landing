NAME:=$(shell basename `git rev-parse --show-toplevel`)
HASH:=$(shell git rev-parse --verify --short HEAD)

deploy: build
	gcloud app deploy

up: build
	npm start

build:
	npm run build

init:
	npm install
