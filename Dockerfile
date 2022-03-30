FROM golang:1.16-alpine
WORKDIR /go/src/app
COPY . .
RUN go get && go install && go build
CMD ["./microservices-boilerplate"]