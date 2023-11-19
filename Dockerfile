FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o msgs ./src/cmd/

EXPOSE 8080

CMD [ "./msgs" ]

