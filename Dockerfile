FROM golang:1.19.2-alpine

WORKDIR /go/src/app

COPY go.mod go.sum ./
COPY cmd/ ./cmd
COPY pkg/ ./pkg

RUN go mod download && go mod verify

RUN go build -o go-todo-app cmd/go-todo-app/main.go
COPY wait-for-mysql.sh ./

EXPOSE 8080

CMD ["./go-todo-app"]
