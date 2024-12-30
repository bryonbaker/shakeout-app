# Makefile for building shakeout-app container

IMAGE_NAME := quay.io/bryonbaker/shakeout-app
TAG := latest

.PHONY: all build push clean

all: build push

build:
	podman build -t $(IMAGE_NAME):$(TAG) .

push:
	podman push $(IMAGE_NAME):$(TAG)

clean:
	podman rmi -f $(IMAGE_NAME):$(TAG) || true
