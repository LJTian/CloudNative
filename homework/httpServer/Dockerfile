FROM golang:1.16-alpine AS build

WORKDIR /go/src/httpServer/

COPY . /go/src/httpServer/

RUN go build -0 /bin/httpServerExec

FROM scratch

COPY --from=build /bin/httpServerExec /bin/httpServer

ENTRYPOINT ["/bin/httpServer"]

