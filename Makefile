.PHONY: default
default: build

all: clean get-deps build test

version := "0.1.0"

build:
	mkdir -p bin
	cd comps/golang && go build -o ../../bin/service-sonar main.go

test: build
	#cd comps/golang && go test -short -coverprofile=../../bin/cov.out
	#cd comps/golang && go tool cover -func=../../bin/cov.out
	#cat bin/cov.out
	#go test -short -coverprofile=bin/cov.out ./...
	#go tool cover -func=bin/cov.out


clean:
	rm -rf ./bin

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube