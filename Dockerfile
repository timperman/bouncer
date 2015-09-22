FROM golang:1.5.1

EXPOSE 8080
COPY . /go/src/github.com/timperman/bouncer

RUN go get -d -v ./...

RUN go test ./... && go install ./...

ENTRYPOINT [ "/go/bin/bouncer" ]
