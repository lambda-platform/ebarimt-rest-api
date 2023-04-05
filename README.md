# REST API for EBarimt PosAPI with Go lang

## Go package for EBarimt PosAPI

go get github.com/lambda-platform/ebarimt

## Powered By Developer team of Lambda Platform.

swag init --parseDependency -o docs/mn
swag init --parseDependency -o docs/en

### [Main Package](https://github.com/lambda-platform/ebarimt)


docker build --platform linux/amd64 --build-arg ARCH=x64 -t ebarimt .

docker run --platform linux/amd64 -p 3000:3000 --name=ebarimt ebarimt