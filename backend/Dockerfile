FROM golang:1.14

EXPOSE 8080
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o salary-go-api.bin main.go

CMD ["./salary-go-api.bin"]