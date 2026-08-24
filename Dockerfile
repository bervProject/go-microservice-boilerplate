FROM golang:1.27-alpine
WORKDIR /go/src/app
COPY . .
RUN go get && go install && go build
CMD ["./go-microservice-boilerplate"]