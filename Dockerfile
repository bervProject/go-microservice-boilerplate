FROM golang:1.24.5-alpine
WORKDIR /go/src/app
COPY . .
RUN go get && go install && go build
CMD ["./go-microservice-boilerplate"]