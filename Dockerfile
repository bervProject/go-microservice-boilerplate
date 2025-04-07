FROM golang:1.24.2-alpine
WORKDIR /go/src/app
COPY . .
RUN go get && go install && go build
CMD ["./go-microservice-boilerplate"]