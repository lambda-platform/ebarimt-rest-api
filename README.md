# REST API for EBarimt PosAPI with Go lang

## Go package for EBarimt PosAPI

go get github.com/lambda-platform/ebarimt

## Powered By Developer team of Lambda Platform.

swag init --parseDependency -o docs/mn
swag init --parseDependency -o docs/en

### [Main Package https://github.com/lambda-platform/ebarimt ](https://github.com/lambda-platform/ebarimt)



Docker build image command
docker build --platform linux/amd64 --build-arg ARCH=x64 -t ebarimt .


check so in ubuntu example
nm -D /home/ebarimtuser/app/libPosAPI.so | grep getInformation
nm -D /usr/lib/libPosAPI.so | grep getInformation