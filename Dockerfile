FROM golang:1.19.2-alpine
WORKDIR /go/src/app
COPY . .
RUN go get && go install && go build
CMD ["./go-microservice-boilerplate"]