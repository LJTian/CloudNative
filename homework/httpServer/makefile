tag=v1.0
UserName="LJTian"
PathName="homeWork"
ExecName="httpServer"

TagName=${UserName}/${PathName}/${ExecName}:${tag}

build:
	echo "building httpserver binary"
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main src/main.go

# DockerFile 文件中已经进行编译了，不需要依赖build进行编译
release:
	echo "build -t ${TagName} ."
	docker build -t ${TagName} .

push: release
	echo "pushing ${tagName}"
	docker push ${TagName}